package controllers

import (
	"fmt"
	"net/http"
)

func IndexController(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "REST API running ...")
}
