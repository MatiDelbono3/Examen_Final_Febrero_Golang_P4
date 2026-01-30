package dtos

type ListarPaginadoResponse struct {
	Total int              `json:"total"`
	Items []RecetaResponse `json:"Items"`
}
