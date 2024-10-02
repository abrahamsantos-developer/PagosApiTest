package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"myPagosApp/internal/models"
	"myPagosApp/internal/services"
	"net/http"
)

// maneja solicitudes HTTP de transactions
type TransactionHandler struct {
	service *services.TransactionService
}

// crea nuevo handler para transactions
func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// valida el body en Swagger(example)
type SwaggerTransactionRequest struct {
	MerchantID uuid.UUID `json:"merchant_id" example:"d290f1ee-6c54-4b01-90e6-d701748f0851"`
	Amount     float64   `json:"amount" example:"100.50"`
}

// struct personalizado(error)
type TransactionErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
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
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{
			Error:   "Error de validacion",
			Message: "El cuerpo de la solicitud no es valido: " + err.Error(),
		})
		return
	}

	// valida que MerchantID sea UUID
	if _, err := uuid.Parse(transaction.MerchantID.String()); err != nil {
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{
			Error:   "UUID con formato invalido",
			Message: "El MerchantID proporcionado no es un UUID valido",
		})
		return
	}

	// valida que ammount no sea negativo o cero
	if transaction.Amount <= 0 {
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{
			Error:   "Monto invalido",
			Message: "El monto de la transaccion debe ser mayor a 0",
		})
		return
	}

	if err := h.service.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{
			Error:   "Error al crear la transacción",
			Message: err.Error(),
		})
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
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{
			Error:   "UUID con formato invalido",
			Message: "El MerchantID proporcionado no es un UUID valido",
		})
		return
	}

	transactions, err := h.service.GetTransactionsByMerchantID(merchantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TransactionErrorResponse{
			Error:   "Error al obtener las transacciones",
			Message: "Hubo un problema al intentar obtener las transacciones del comercio: " + err.Error(),
		})
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
		c.JSON(http.StatusBadRequest, TransactionErrorResponse{
			Error:   "UUID con formato invalido",
			Message: "La ID proporcionada no es un UUID valido",
		})
		return
	}

	transaction, err := h.service.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, TransactionErrorResponse{
			Error:   "Transacción no encontrada",
			Message: "No se encontró una transacción con el ID proporcionado",
		})
		return
	}
	c.JSON(http.StatusOK, transaction)
}
