package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest_api/app/models"
	"rest_api/app/services"
)

func PostTraining(response http.ResponseWriter, request *http.Request) {
	var parameters models.Parameters
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Fprintf(response, "Parametros inválidos")
	}

	json.Unmarshal(body, &parameters)
	fit_K, fit_accuracy := services.TrainingService(parameters)
	res := models.FitKNN{K: fit_K, Accuracy: fit_accuracy}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode(res)
}
