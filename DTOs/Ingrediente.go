package dtos

type Ingrediente struct {
	Nombre   string  `json:"nombre"`
	Cantidad float64 `json:"cantidad"`
	Unidad   string  `json:"unidad"`
}
