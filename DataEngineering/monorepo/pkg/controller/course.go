package controller

import (
	"encoding/json"
	"net/http"

	"api_thienhang_com/pkg/entity"
	"api_thienhang_com/pkg/utils"
)

func (c *Controller) createCourse(w http.ResponseWriter, r *http.Request) {
	// Get account information from google
	account, err := c.getAccount(r)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusForbidden, err.Error())
		return
	}

	var data entity.Course
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// creatorID := primitive.ObjectID(account.UID)
	// data.Lecturer = []primitive.ObjectID{creatorID}
	userInfo, err := c.service.CreateCourse(account, data)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, userInfo)
}

func (c *Controller) updateCourse(w http.ResponseWriter, r *http.Request) {
	var data entity.Course
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	userInfo, err := c.service.UpdateCourse(data)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, userInfo)
}

func (c *Controller) getCourse(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("id")

	userInfo, err := c.service.GetCourse(uuid)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, userInfo)
}

func (c *Controller) createLession(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("id")

	var data entity.Lession
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	err := c.service.CreateLession(uuid, &data)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, nil)
}
