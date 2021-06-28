package models

type FitKNN struct {
	K        int     `json:"k"`
	Accuracy float64 `json:"accuracy"`
}
