package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"myPagosApp/internal/services"
	"net/http"
)

// ProfitHandler maneja las solicitudes relacionadas con ganancias
type ProfitHandler struct {
	service *services.ProfitService
}

// NewProfitHandler crea un nuevo handler para ganancias
func NewProfitHandler(service *services.ProfitService) *ProfitHandler {
	return &ProfitHandler{service: service}
}

// Estructura de respuesta para errores de ganancias
type ProfitErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Obtener las ganancias totales
// @Description Obtiene las ganancias de todas las transacciones registradas.
// @Tags Ganancias
// @Produce json
// @Success 200 {object} map[string]float64
// @Failure 500 {object} ProfitErrorResponse
// @Router /profits [get]
func (h *ProfitHandler) GetTotalProfitsHandler(c *gin.Context) {
	totalProfits, err := h.service.GetTotalProfits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ProfitErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_profits": totalProfits})
}

// @Summary Obtener las ganancias por comercio
// @Description Obtiene las ganancias de un comercio específico mediante su ID.
// @Tags Ganancias
// @Produce json
// @Param merchant_id path string true "ID del Comercio"
// @Success 200 {object} map[string]float64
// @Failure 400 {object} ProfitErrorResponse
// @Failure 500 {object} ProfitErrorResponse
// @Router /profits/merchant/{merchant_id} [get]
func (h *ProfitHandler) GetProfitsByMerchantIDHandler(c *gin.Context) {
	idParam := c.Param("merchant_id")
	merchantID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ProfitErrorResponse{Error: "ID inválido"})
		return
	}

	profits, err := h.service.GetProfitsByMerchantID(merchantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ProfitErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"merchant_profits": profits})
}
