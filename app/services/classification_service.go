package services

import (
	"rest_api/app/knn"
	"rest_api/app/models"
	"rest_api/app/utils"
)

func ClassificationService(K int, obj models.ClassifyData) (class string) {
	lines, err := utils.GetFileByUrl(utils.Url_data)

	if err != nil {
		print(err)
	}

	sexo := float64(obj.Sexo)
	edad := float64(obj.Edad)
	etnia := float64(obj.Etnia)
	nivel_educativo := float64(obj.NivelEducativo)
	ultimo_cargo := float64(obj.UltimoCargo)
	frecuencia_pago := float64(obj.FrecuenciaPago)
	ingreso_monetario := obj.IngresoMonetario
	seguro_salud := float64(obj.SeguroSalud)

	persona_to_classify := utils.PersonaEncuestada{
		Data:  []float64{sexo, edad, etnia, nivel_educativo, ultimo_cargo, frecuencia_pago, ingreso_monetario, seguro_salud},
		Class: "",
	}

	personas := CleanData(lines)
	class = knn.ClassifyClass(persona_to_classify, personas, K)
	return class
}
