package dtos

import (
	models "examen_final_febrero_golang_P4/Models"
	"time"
)

type RecetaResponse struct {
	Id            string        `json:"id"`
	Nombre        string        `json:"nombre"`
	Categoria     string        `json:"categoria"`
	Ingredientes  []Ingrediente `json:"ingredientes"`
	FechaCreacion time.Time     `json:"fechaCreacion"`
	IdUsuario     string        `json:"idUsuario"`
}

func NewFood(Receta models.Receta) *RecetaResponse {
	return &RecetaResponse{
		Nombre:        Receta.Nombre,
		Categoria:     Receta.Categoria,
		Ingredientes:  Receta.Ingredientes,
		FechaCreacion: Receta.FechaCreacion,
	}
}
