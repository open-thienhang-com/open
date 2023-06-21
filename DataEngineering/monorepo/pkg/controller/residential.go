package controller

import (
	"encoding/json"
	"net/http"

	"api_thienhang_com/pkg/entity"
	"api_thienhang_com/pkg/utils"
)

// ShowAccount godoc
// @Summary      Lấy thông tin của user, nếu không có thì đồng bộ từ firebase
// @Description  get string by ID
// @Tags         Residential
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Address
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /residential [post]
func (c *Controller) createResidential(w http.ResponseWriter, r *http.Request) {
	// Get account information from google
	account, err := c.getAccount(r)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusForbidden, err.Error())
		return
	}

	var data entity.Address
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// creatorID := primitive.ObjectID(account.UID)
	// data.Lecturer = []primitive.ObjectID{creatorID}
	userInfo, err := c.service.AddResidential(account, &data)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, userInfo)
}

// ShowAccount godoc
// @Summary      Cập nhật thông tin cho người dùng
// @Description  get string by ID
// @Tags         Residential
// @Param input body entity.Address true "Add InputCreateListener"
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Address
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /residential [put]
func (c *Controller) updateResidential(w http.ResponseWriter, r *http.Request) {
	account, err := c.getAccount(r)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusForbidden, err.Error())
		return
	}
	var data entity.Address
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	userInfo, err := c.service.UpdateResidential(account, &data)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, userInfo)
}

// ShowAccount godoc
// @Summary
// @Description  get string by ID
// @Tags         Residential
// @Param   province      query     string     false  "string valid"       minlength(0)  maxlength(30)
// @Param   district      query     string     false  "string valid"       minlength(0)  maxlength(30)
// @Param   ward      query     string     false  "string valid"       minlength(0)  maxlength(30)
// @Param   building      query     string     false  "string valid"       minlength(0)  maxlength(30)
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Province
// @Failure      400
// @Failure      404
// @Failure      500
// @Router        /residential [get]
func (c *Controller) getResidential(w http.ResponseWriter, r *http.Request) {
	province := r.URL.Query().Get("province")
	if province == "all" || province == "" {
		userInfo, err := c.service.GetProvinces(nil)
		if err != nil {
			utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.ResponseWithJson(w, http.StatusOK, userInfo)
		return
	}

	district := r.URL.Query().Get("district")
	if district == "all" || district == "" {
		userInfo, err := c.service.GetDistricts(nil, province)
		if err != nil {
			utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.ResponseWithJson(w, http.StatusOK, userInfo)
		return
	}

	ward := r.URL.Query().Get("ward")
	if ward == "all" || ward == "" {
		userInfo, err := c.service.GetWards(nil, province, district)
		if err != nil {
			utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.ResponseWithJson(w, http.StatusOK, userInfo)
		return
	}

	building := r.URL.Query().Get("building")
	if building == "all" || building == "" {
		userInfo, err := c.service.GetBuildings(nil, province, district, ward)
		if err != nil {
			utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.ResponseWithJson(w, http.StatusOK, userInfo)
		return
	}
	// //
	utils.ResponseWithJson(w, http.StatusBadRequest, "Vui lòng kiểm tra lại điều kiện lọc")
	return
}

// ShowAccount godoc
// @Summary
// @Description  get string by ID
// @Tags         Residential
// @Param uuid   path string true "UUID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Address
// @Failure      400
// @Failure      404
// @Failure      500
// @Router        /residential/address [get]
func (c *Controller) getAddress(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")

	userInfo, err := c.service.GetAddress(nil, uuid)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, userInfo)
}
