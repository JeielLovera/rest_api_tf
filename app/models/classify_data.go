package models

type ClassifyData struct {
	K                int     `json:"k"`
	Sexo             int     `json:"sexo"`
	Edad             int     `json:"edad"`
	Etnia            int     `json:"etnia"`
	NivelEducativo   int     `json:"nivelEducativo"`
	UltimoCargo      int     `json:"ultimoCargo"`
	FrecuenciaPago   int     `json:"frecuenciaPago"`
	IngresoMonetario float64 `json:"ingresoMonetario"`
	SeguroSalud      int     `json:"seguroSalud"`
	Class            string  `json:"class"`
}
