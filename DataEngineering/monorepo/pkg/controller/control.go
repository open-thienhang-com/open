package controller

// @title Tài liệu mô tả cho module chat
// @version 1.0
// @description

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description

// @BasePath /api/v1
// @schemes http https
// @query.collection.format multi

// @contact.name API Support 122012 - Hàng Tuấn Thiên
// @contact.url
// @x-extension-openapi {"example": "value on a json format"}

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// "github.com/sirupsen/logrus"

	"net/http"
	"net/http/pprof"
	// "fmt"
	// "path/filepath"
	servicepk "api_thienhang_com/pkg/service"
)

type Controller struct {
	router  *mux.Router
	service servicepk.IService
}

func Init() {

}
func New(router *mux.Router) *Controller {
	router.Use(mux.CORSMethodMiddleware(router))
	s := servicepk.Init()
	s.InitResidential()
	return &Controller{
		router:  router,
		service: s,
	}
}

func (c *Controller) ServeHTTP() {
	// logrus.Info("Xin cháo")
	// path, err := filepath.Abs(".")
	// if err != nil {
    //     // if we failed to get the absolute path respond with a 400 bad request
    //     // and stop
	// 	fmt.Println(err)
	// }

    // //
	// // staticPaths := map[string]string{
	// // 	"styles":           staticDirectory + "/styles/",
	// // 	"bower_components": staticDirectory + "/bower_components/",
	// // 	"images":           staticDirectory + "/images/",
	// // 	"scripts":          staticDirectory + "/scripts/",
	// // }
	// // for pathName, pathValue := range staticPaths {
	// // 	pathPrefix := "/" + pathName + "/"
	// // 	router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
	// // 		http.FileServer(http.Dir(pathValue))))
	// // }

	// //
	// fmt.Println(path+"/public/")
	// c.router.PathPrefix("/public").Handler(http.FileServer(http.Dir(path+"/public/")))
	
	//
	c.router.HandleFunc("/debug", pprof.Index)
	c.router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	c.router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	c.router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	c.router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	c.router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	c.router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	c.router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	c.router.Handle("/debug/pprof/block", pprof.Handler("block"))

	c.router.Handle("/metrics", promhttp.Handler())


	//
	apiv1 := c.router.PathPrefix("/api/v1").Subrouter()
	// ****************************************************************
	// *** PORTAL ***
	apiv1.HandleFunc("/notify/email", c.sendEmail).Methods(http.MethodPost)

	// ****************************************************************
	// *** USER MANAGEMENT ***
	apiv1.HandleFunc("/user", c.checkUser).Methods(http.MethodPost)
	apiv1.HandleFunc("/user", c.updateUser).Methods(http.MethodPut)

	// *** RESIDENTIAL ***
	apiv1.HandleFunc("/residential", c.createResidential).Methods(http.MethodPost)
	apiv1.HandleFunc("/residential", c.updateResidential).Methods(http.MethodPut)

	apiv1.HandleFunc("/residential", c.getResidential).Methods(http.MethodGet)

	apiv1.HandleFunc("/residential/address", c.getAddress).Methods(http.MethodGet)

	// *** EDUCATION ***
	// apiv1.HandleFunc("/user/education", c.AddEducation).Methods(http.MethodPost)
	// apiv1.HandleFunc("/user/education", c.UpdateEducation).Methods(http.MethodPut)
	// apiv1.HandleFunc("/user/education", c.DeleteEducation).Methods(http.MethodDelete)

	// // *** EXPERIENCE ***
	// apiv1.HandleFunc("/user/experience", c.AddExperience).Methods(http.MethodPost)
	// apiv1.HandleFunc("/user/experience", c.UpdateExperience).Methods(http.MethodPut)
	// apiv1.HandleFunc("/user/experience", c.DeleteExperience).Methods(http.MethodDelete)

	// // *** SKILL ***
	// apiv1.HandleFunc("/user/skill", c.AddSkill).Methods(http.MethodPost)
	// apiv1.HandleFunc("/user/skill", c.UpdateSkill).Methods(http.MethodPut)
	// apiv1.HandleFunc("/user/skill", c.DeleteSkill).Methods(http.MethodDelete)

	// // *** REFERENCE ***
	// apiv1.HandleFunc("/user/reference", c.AddReference).Methods(http.MethodPost)
	// apiv1.HandleFunc("/user/reference", c.UpdateReference).Methods(http.MethodPut)
	// apiv1.HandleFunc("/user/reference", c.DeleteReference).Methods(http.MethodDelete)

	// // *** ADWARD ***
	// apiv1.HandleFunc("/user/award", c.AddAward).Methods(http.MethodPost)
	// apiv1.HandleFunc("/user/award", c.UpdateAward).Methods(http.MethodPut)
	// apiv1.HandleFunc("/user/award", c.DeleteAward).Methods(http.MethodDelete)

	// // *** QUALIFICATION ***
	// apiv1.HandleFunc("/user/qualification", c.AddQualification).Methods(http.MethodPost)
	// apiv1.HandleFunc("/user/qualification", c.UpdateQualification).Methods(http.MethodPut)
	// apiv1.HandleFunc("/user/qualification", c.DeleteQualification).Methods(http.MethodDelete)

	// ****************************************************************
	// *** COURSE ***
	apiv1.HandleFunc("/course", c.getCourse).Methods(http.MethodGet)
	apiv1.HandleFunc("/courses", c.createCourse).Methods(http.MethodGet)
	apiv1.HandleFunc("/course", c.createCourse).Methods(http.MethodPost)
	apiv1.HandleFunc("/course", c.updateCourse).Methods(http.MethodPut)
	apiv1.HandleFunc("/course", c.updateCourse).Methods(http.MethodDelete)

	// *** LESSION ***
	apiv1.HandleFunc("/lession", c.createCourse).Methods(http.MethodGet)
	apiv1.HandleFunc("/lessions", c.createCourse).Methods(http.MethodGet)
	apiv1.HandleFunc("/lession", c.createLession).Methods(http.MethodPost)
	apiv1.HandleFunc("/lession", c.createCourse).Methods(http.MethodPost)
	apiv1.HandleFunc("/lession", c.updateCourse).Methods(http.MethodPut)

	c.printEndpoints()

}
