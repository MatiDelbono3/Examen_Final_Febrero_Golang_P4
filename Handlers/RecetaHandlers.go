package handlers

import (
	Services "examen_final_febrero_golang_P4/Service"
	Dtos "examen_final_febrero_golang_P4/dtos"
	"examen_final_febrero_golang_P4/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type RecetaHandler struct {
	service Services.RecetaServiceInterface
}

func NewRecetaHandler(service Services.RecetaServiceInterface) *RecetaHandler {
	return &RecetaHandler{
		service: service,
	}
}

func (handler *RecetaHandler) Crear(c *gin.Context) {
	var request Dtos.RecetaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receta, err := handler.service.Crear(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, receta)
}
func (handlers *RecetaHandler) ListarPaginado(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit inválido"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "offset inválido"})
		return
	}
	resp, err := handlers.service.ListarPaginado(limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (handlers *RecetaHandler) FiltrarRecetasPorCategoria(c *gin.Context) {
	categoria:=c.Param("categoria")
	resp, err := h.service.FiltrarRecetasPorNombre(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (handlers *RecetaHandler) FiltrarRecetasPorNombre(c *gin.Context) {
	nombre:=c.Param("nombre")
	resp, err := h.service.FiltrarRecetasPorNombre(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (handlers *RecetaHandler) FiltrarRecetasPorID(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.service.FiltrarRecetasPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}