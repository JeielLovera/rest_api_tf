package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest_api/app/models"
	"rest_api/app/services"
)

func PostClassification(response http.ResponseWriter, request *http.Request) {
	var classify_data models.ClassifyData
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Fprintf(response, "Parametros inválidos")
	}

	json.Unmarshal(body, &classify_data)
	class := services.ClassificationService(classify_data.K, classify_data)
	classify_data.Class = class

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(classify_data)
}
