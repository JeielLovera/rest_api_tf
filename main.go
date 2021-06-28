package main

import (
	"net/http"
	"rest_api/app"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := app.New()
	http.HandleFunc("/", app.Router.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
