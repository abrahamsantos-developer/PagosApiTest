package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"myPagosApp/internal/models"
	"myPagosApp/internal/services"
	"net/http"
)

// maneja solicitudes HTTP relacionadas de merchant
type MerchantHandler struct {
	service *services.MerchantService
}

// creanuevo handlerpara merchant
func NewMerchantHandler(service *services.MerchantService) *MerchantHandler {
	return &MerchantHandler{service: service}
}

type SwaggerMerchantRequest struct {
	Name       string `json:"name" example:"comercio123"`
	Commission uint   `json:"commission" example:"15"`
}

// struct de respuesta
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// @Summary Crear un comercio
// @Description Crear un nuevo comercio en el sistema.
// @Tags Comercios
// @Accept  json
// @Produce  json
// @Param merchant body handlers.SwaggerMerchantRequest true "Comercio a crear"
// @Success 200 {object} models.Merchant
// @Failure 400 {object} ErrorResponse
// @Router /merchants [post]
func (h *MerchantHandler) CreateMerchantHandler(c *gin.Context) {
	var merchant models.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Error de validacion",
			Message: "El cuerpo de la solicitud no es valido: " + err.Error(),
		})
		return
	}

	// valida que el name no este vacio
	if merchant.Name == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Error de validacion",
			Message: "El nombre del comercio no puede estar vacio",
		})
		return
	}

	if err := h.service.CreateMerchant(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Error al crear el comercio",
			Message: err.Error(),
		})
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

// @Summary Obtener un comercio por ID
// @Description Obtener los detalles de un comercio específico mediante su ID.
// @Tags Comercios
// @Produce json
// @Param id path string true "ID del Comercio"
// @Success 200 {object} models.Merchant
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /merchants/{id} [get]
func (h *MerchantHandler) GetMerchantByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "UUID con formato invalido",
			Message: "La ID proporcionada no es un UUID valido",
		})
		return
	}

	merchant, err := h.service.GetMerchantByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "Comercio no encontrado",
			Message: "No se encontró un comercio con el ID proporcionado",
		})
		return
	}
	c.JSON(http.StatusOK, merchant)
}

// @Summary Actualizar un comercio
// @Description Actualizar los detalles de un comercio existente.
// @Tags Comercios
// @Accept  json
// @Produce  json
// @Param id path string true "ID del Comercio"
// @Param merchant body handlers.SwaggerMerchantRequest true "Datos actualizados del comercio"
// @Success 200 {object} models.Merchant
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /merchants/{id} [put]
func (h *MerchantHandler) UpdateMerchantHandler(c *gin.Context) {
	// Obtiene ID desde el param de la ruta
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "UUID con formato inválido",
			Message: "La ID proporcionada no es un UUID valido",
		})
		return
	}

	// obtener datos de update del cuerpo del request
	var updatedMerchant models.Merchant
	if err := c.ShouldBindJSON(&updatedMerchant); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Error de validación",
			Message: "El cuerpo de la solicitud no es válido: " + err.Error(),
		})
		return
	}

	// Asegurarse de no resetear el ID
	updatedMerchant.ID = id

	// try update merchant
	if err := h.service.UpdateMerchant(id, &updatedMerchant); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Error al actualizar el comercio",
			Message: err.Error(),
		})
		return
	}

	// Volver a obtener el comercio actualizado desde la base de datos (con las transacciones)
	merchant, err := h.service.GetMerchantByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Error al obtener el comercio actualizado",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, merchant)
}
