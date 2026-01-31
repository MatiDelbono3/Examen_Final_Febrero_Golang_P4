package dtos

import (
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
