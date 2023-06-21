package controller

import (
	"github.com/gorilla/mux"
	"fmt"
)

func (c *Controller) printEndpoints() {
	// c.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 	path, err1 := route.GetPathTemplate()
	// 	if err1 != nil {
	// 		return nil
	// 	}
	// 	methods, err2 := route.GetMethods()
	// 	if err2 != nil {
	// 		return nil
	// 	}
	// 	fmt.Printf("%v %s\n %v %v", methods, path, err1, err2)
	// 	return nil
	// })
	// DEBUG
	c.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})
}