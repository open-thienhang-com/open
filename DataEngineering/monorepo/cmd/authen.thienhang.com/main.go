package main

import (
	"net/http"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"api_thienhang_com/pkg/controller"

	_ "api_thienhang_com/cmd/authen.thienhang.com/docs"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func init() {
	runtime.GOMAXPROCS(1)
	err := godotenv.Load("config.env")
	if err != nil {
		log.Error("Error loading .env file")
	}
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{})
	//
	//
}

// @title AUTHENTICATION OPEN API - thienhang.com
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService https://thienhang.com

// @contact.name API Support
// @contact.url http://thienhang.com
// @contact.email me@thienhang.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @name OpenKey
func main() {
	m := mux.NewRouter()
	m.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("docs/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
		httpSwagger.BeforeScript(`const UrlMutatorPlugin = (system) => ({
		rootInjects: {
			setScheme: (scheme) => {
			const jsonSpec = system.getState().toJSON().spec.json;
			const schemes = Array.isArray(scheme) ? scheme : [scheme];
			const newJsonSpec = Object.assign({}, jsonSpec, { schemes });

			return system.specActions.updateJsonSpec(newJsonSpec);
			},
			setHost: (host) => {
			const jsonSpec = system.getState().toJSON().spec.json;
			const newJsonSpec = Object.assign({}, jsonSpec, { host });

			return system.specActions.updateJsonSpec(newJsonSpec);
			},
			setBasePath: (basePath) => {
			const jsonSpec = system.getState().toJSON().spec.json;
			const newJsonSpec = Object.assign({}, jsonSpec, { basePath });

			return system.specActions.updateJsonSpec(newJsonSpec);
			}
		}
		
		});`),
		httpSwagger.Plugins([]string{"UrlMutatorPlugin"}),
		httpSwagger.UIConfig(map[string]string{
			"showExtensions":        "true",
			"onComplete":            `() => { window.ui.setBasePath('/api/v1'); }`,
			"defaultModelRendering": `"model"`,
		}),
		// 		httpSwagger.UIConfig(map[string]string{
		// 			"onComplete": fmt.Sprintf(`() => {
		//     window.ui.setScheme('%s');
		//     window.ui.setHost('%s');
		//     window.ui.setBasePath('%s');
		//   }`, uri.Scheme, uri.Host, uri.Path),
		// 		}),
	)).Methods(http.MethodGet)
	//
	controller := controller.New(m)

	controller.ServeHTTP()

	log.Warn("Hello")
	// admin.SendEmail()
	// ip, _ := getIP()
	// admin.SendLineNotify("🎉🎉🎉 AUTHEN: " + ip + " |✅  " + time.Now().GoString())
	//http://localhost:9001/docs/index.html
	router := handlers.CORS(handlers.AllowedHeaders(
		[]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))(m)
	log.Fatal(http.ListenAndServe(":9001", router))

}
