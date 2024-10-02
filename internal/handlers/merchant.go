package handlers

import (
    "myPagosApp/internal/models"
    "myPagosApp/internal/services"
    "github.com/gin-gonic/gin"
    "net/http"
	"github.com/google/uuid"
)

// MerchantHandler maneja las solicitudes HTTP relacionadas con comercios
type MerchantHandler struct {
    service *services.MerchantService
}

// NewMerchantHandler crea un nuevo handler para comercios
func NewMerchantHandler(service *services.MerchantService) *MerchantHandler {
    return &MerchantHandler{service: service}
}

// CreateMerchantHandler maneja la creación de un comercio
func (h *MerchantHandler) CreateMerchantHandler(c *gin.Context) {
    var merchant models.Merchant
    if err := c.ShouldBindJSON(&merchant); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.CreateMerchant(&merchant); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, merchant)
}

// GetAllMerchantsHandler maneja la obtención de todos los comercios
func (h *MerchantHandler) GetAllMerchantsHandler(c *gin.Context) {
    merchants, err := h.service.GetAllMerchants()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, merchants)
}


// UpdateMerchantHandler maneja la actualización de un comercio
func (h *MerchantHandler) UpdateMerchantHandler(c *gin.Context) {
    // Obtener el ID (UUID) desde el parámetro de la ruta
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    // Obtener los nuevos datos del comercio desde el cuerpo de la solicitud
    var updatedMerchant models.Merchant
    if err := c.ShouldBindJSON(&updatedMerchant); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Actualizar el comercio
    if err := h.service.UpdateMerchant(id, &updatedMerchant); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedMerchant)
}

// // UpdateMerchantHandler maneja la actualización de un comercio
// func (h *MerchantHandler) UpdateMerchantHandler(c *gin.Context) {
//     var merchant models.Merchant
//     if err := c.ShouldBindJSON(&merchant); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }
//     if err := h.service.UpdateMerchant(&merchant); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }
//     c.JSON(http.StatusOK, merchant)
// }
