package services

import (
	"fmt"
    "myPagosApp/internal/models"
    "myPagosApp/internal/repositories"
)

// MerchantService define la lógica de negocio para los comercios
type MerchantService struct {
    repository *repositories.MerchantRepository
}

// NewMerchantService crea un nuevo servicio de comercios
func NewMerchantService(repo *repositories.MerchantRepository) *MerchantService {
    return &MerchantService{repository: repo}
}

// CreateMerchant valida y guarda un nuevo comercio
func (s *MerchantService) CreateMerchant(merchant *models.Merchant) error {
    if merchant.Commission < 1 || merchant.Commission > 100 {
        return fmt.Errorf("la comisión debe estar entre 1 y 100")
    }
    return s.repository.CreateMerchant(merchant)
}

// GetAllMerchants obtiene todos los comercios
func (s *MerchantService) GetAllMerchants() ([]models.Merchant, error) {
    return s.repository.GetAllMerchants()
}

// UpdateMerchant actualiza un comercio existente
func (s *MerchantService) UpdateMerchant(merchant *models.Merchant) error {
    return s.repository.UpdateMerchant(merchant)
}
