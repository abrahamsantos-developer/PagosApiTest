package services

import (
	"fmt"
	"github.com/google/uuid"
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
	if merchant.Commission < 1.0 || merchant.Commission > 100.0 {
		return fmt.Errorf("la comisión debe estar entre 1.0 y 100.0")
	}
	return s.repository.CreateMerchant(merchant)
}

// GetAllMerchants obtiene todos los comercios
func (s *MerchantService) GetAllMerchants() ([]models.Merchant, error) {
	return s.repository.GetAllMerchants()
}

// GetMerchantByID busca un comercio por su UUID
func (s *MerchantService) GetMerchantByID(id uuid.UUID) (*models.Merchant, error) {
	return s.repository.GetMerchantByID(id)
}

// UpdateMerchant actualiza un comercio existente por su UUID
func (s *MerchantService) UpdateMerchant(id uuid.UUID, merchant *models.Merchant) error {
	// Buscar si el comercio existe antes de actualizar
	existingMerchant, err := s.repository.GetMerchantByID(id)
	if err != nil {
		return fmt.Errorf("comercio no encontrado: %w", err)
	}

	// Validar la nueva comision
	if merchant.Commission < 1.0 || merchant.Commission > 100.0 {
		return fmt.Errorf("la comisión debe estar entre 1.0 y 100.0")
	}
	// Actualiza la info
	existingMerchant.Name = merchant.Name
	existingMerchant.Commission = merchant.Commission

	return s.repository.UpdateMerchant(id, existingMerchant)
}
