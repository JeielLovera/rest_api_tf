package app

import (
	"net/http"
	"rest_api/app/controllers"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	app.CORS()
	app.Routes()
	return app
}

func (app *App) Routes() {
	app.Router.HandleFunc("/", controllers.IndexController).Methods("GET")
	app.Router.HandleFunc("/api/training", controllers.PostTraining).Methods("POST")
	app.Router.HandleFunc("/api/classification", controllers.PostClassification).Methods("POST")
}

func (app *App) CORS() {
	app.Router.PathPrefix("/").HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	app.Router.Use(OptionsCors)
}

func OptionsCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			response.Header().Set("Access-Control-Allow-Origin", "*")
			response.Header().Set("Access-Control-Allow-Credentials", "true")
			response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			response.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(response, request)
		})
}
