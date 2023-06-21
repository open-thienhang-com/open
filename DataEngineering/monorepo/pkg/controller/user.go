package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"api_thienhang_com/pkg/admin"
	"api_thienhang_com/pkg/entity"
	"api_thienhang_com/pkg/utils"
	"github.com/sirupsen/logrus"
)

// Lấy tài khoản google bằng token
func (c *Controller) getAccount(r *http.Request) (account *entity.Account, err error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, errors.New("wrong token")
	}
	account, err = c.service.GetAccount(token)
	if err != nil {
		return nil, errors.New("can't authorization")
	}
	return account, nil
}

// ShowAccount godoc
// @Summary      Lấy thông tin của user, nếu không có thì đồng bộ từ firebase
// @Description  get string by ID
// @Tags         Users
// @Param Authorization header string true "With the bearer started"
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.User
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user [post]
func (c *Controller) checkUser(w http.ResponseWriter, r *http.Request) {
	// Get account information from google
	acc, err := c.getAccount(r)
	logrus.Info(acc.PhoneNumber)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusForbidden, err.Error())
		return
	}
	admin.SendNotification("🚀  Xin chào anh chị >" + fmt.Sprint(acc.Email, acc.PhoneNumber))
	userInfo, err := c.service.CheckUser(acc)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, userInfo)
}

// ShowAccount godoc
// @Summary      Cập nhật thông tin cho người dùng
// @Description  get string by ID
// @Tags         Users
// @Param Authorization header string true "Điền token firebase"
// @Param input body entity.User true "Add InputCreateListener"
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.User
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user [put]
func (c *Controller) updateUser(w http.ResponseWriter, r *http.Request) {
	// Get account information from google
	account, err := c.getAccount(r)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusForbidden, err.Error())
		return
	}

	// Get Update information
	var data entity.User
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := c.service.UpdateUser(account, &data)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, newUser)
}
