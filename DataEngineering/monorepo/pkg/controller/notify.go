package controller

import (
	
	"net/http"
	"fmt"
	"api_thienhang_com/pkg/admin"
	"api_thienhang_com/pkg/utils"
	
)

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
// @Router       /notify/email [post]
func (c *Controller) sendEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OKKKKKK!")
	admin.SendEmail()
	utils.ResponseWithJson(w, http.StatusOK, "OK")
}