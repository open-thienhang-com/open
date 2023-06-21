package controller

import (
	"github.com/gorilla/mux"

	service "api_thienhang_com/services/demo/service"
)

type Controller struct {
	router  *mux.Router
	service service.IService
}

func New(router *mux.Router) *Controller {
	router.Use(mux.CORSMethodMiddleware(router))
	s := service.Init()
	return &Controller{
		router:  router,
		service: s,
	}
}

func (c *Controller) ServeHTTP() {
	
	
}
