package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"myPagosApp/internal/models"
	"myPagosApp/internal/services"
	"net/http"
)

// TransactionHandler maneja las solicitudes HTTP relacionadas con transacciones
type TransactionHandler struct {
	service *services.TransactionService
}

// NewTransactionHandler crea un nuevo handler para transacciones
func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// SwaggerTransactionRequest se usa para validar el body en Swagger
type SwaggerTransactionRequest struct {
	MerchantID uuid.UUID `json:"merchant_id" example:"d290f1ee-6c54-4b01-90e6-d701748f0851"`
	Amount     float64   `json:"amount" example:"100.50"`
}

// Estructura de respuesta de error
type TransactionErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Crear una transacción
// @Description Crear una nueva transacción para un comercio y calcular el fee basado en la comisión.
// @Tags Transacciones
// @Accept  json
// @Produce  json
// @Param transaction body handlers.SwaggerTransactionRequest true "Transacción a crear"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} TransactionErrorResponse
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransactionHandler(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{Error: err.Error()})
		return
	}
	if err := h.service.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

// @Summary Obtener todas las transacciones
// @Description Obtener una lista de todas las transacciones realizadas.
// @Tags Transacciones
// @Produce  json
// @Success 200 {array} models.Transaction
// @Failure 500 {object} TransactionErrorResponse
// @Router /transactions [get]
func (h *TransactionHandler) GetAllTransactionsHandler(c *gin.Context) {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, TransactionErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// @Summary Obtener todas las transacciones de un comercio
// @Description Obtener todas las transacciones de un comercio específico.
// @Tags Transacciones
// @Produce  json
// @Param merchant_id path string true "ID del Comercio"
// @Success 200 {array} models.Transaction
// @Failure 400 {object} TransactionErrorResponse
// @Router /transactions/merchant/{merchant_id} [get]
func (h *TransactionHandler) GetTransactionsByMerchantIDHandler(c *gin.Context) {
	idParam := c.Param("merchant_id")
	merchantID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{Error: "ID de comercio inválido"})
		return
	}
	transactions, err := h.service.GetTransactionsByMerchantID(merchantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TransactionErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// @Summary Obtener una transacción por ID
// @Description Obtener los detalles de una transacción específica por su ID.
// @Tags Transacciones
// @Produce json
// @Param id path string true "ID de la Transacción"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} TransactionErrorResponse
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransactionByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{Error: "ID inválido"})
		return
	}
	transaction, err := h.service.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, TransactionErrorResponse{Error: "Transacción no encontrada"})
		return
	}
	c.JSON(http.StatusOK, transaction)
}
