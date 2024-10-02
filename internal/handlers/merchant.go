package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"myPagosApp/internal/models"
	"myPagosApp/internal/services"
	"net/http"
)

// MerchantHandler maneja las solicitudes HTTP relacionadas con comercios
type MerchantHandler struct {
	service *services.MerchantService
}

// NewMerchantHandler crea un nuevo handler para comercios
func NewMerchantHandler(service *services.MerchantService) *MerchantHandler {
	return &MerchantHandler{service: service}
}

// Estructuras de respuesta
type ErrorResponse struct {
	Error string `json:"error"`
}

// type MessageResponse struct {
// 	Message string `json:"message"`
// }

// @Summary Crear un comercio
// @Description Crear un nuevo comercio en el sistema.
// @Tags Comercios
// @Accept  json
// @Produce  json
// @Param merchant body models.Merchant true "Comercio a crear"
// @Success 200 {object} models.Merchant
// @Failure 400 {object} ErrorResponse
// @Router /merchants [post]
func (h *MerchantHandler) CreateMerchantHandler(c *gin.Context) {
	var merchant models.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	if err := h.service.CreateMerchant(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, merchant)
}

// @Summary Obtener todos los comercios
// @Description Obtener una lista de todos los comercios en el sistema.
// @Tags Comercios
// @Produce  json
// @Success 200 {array} models.Merchant
// @Failure 500 {object} ErrorResponse
// @Router /merchants [get]
func (h *MerchantHandler) GetAllMerchantsHandler(c *gin.Context) {
	merchants, err := h.service.GetAllMerchants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, merchants)
}

// @Summary Actualizar un comercio
// @Description Actualizar los detalles de un comercio existente.
// @Tags Comercios
// @Accept  json
// @Produce  json
// @Param id path string true "ID del Comercio"
// @Param merchant body models.Merchant true "Datos actualizados del comercio"
// @Success 200 {object} models.Merchant
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /merchants/{id} [put]
func (h *MerchantHandler) UpdateMerchantHandler(c *gin.Context) {
	// Obtener el ID (UUID) desde el parámetro de la ruta
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "ID invalido"})
		return
	}

	// Obtener los nuevos datos del comercio desde el cuerpo de la solicitud
	var updatedMerchant models.Merchant
	if err := c.ShouldBindJSON(&updatedMerchant); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// Actualizar el comercio
	if err := h.service.UpdateMerchant(id, &updatedMerchant); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedMerchant)
}
