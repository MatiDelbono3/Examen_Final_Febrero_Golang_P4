package dtos

type RecetaRequest struct {
	Nombre       string        `json:"nombre"`
	Categoria    string        `json:"categoria"`
	Ingredientes []Ingrediente `json:"ingredientes"`
}
