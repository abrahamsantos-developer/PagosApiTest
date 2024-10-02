package services

import (
	"fmt"
	"github.com/google/uuid"
	"myPagosApp/internal/models"
	"myPagosApp/internal/repositories"
)

// logica para merchants
type MerchantService struct {
	repository *repositories.MerchantRepository
}

// crea nuevo service de merchants
func NewMerchantService(repo *repositories.MerchantRepository) *MerchantService {
	return &MerchantService{repository: repo}
}

// crea y guarda un nuevo merchant
func (s *MerchantService) CreateMerchant(merchant *models.Merchant) error {
	if merchant.Commission < 1.0 || merchant.Commission > 100.0 {
		return fmt.Errorf("la comisión debe estar entre 1.0 y 100.0")
	}
	return s.repository.CreateMerchant(merchant)
}

// obtiene todos los merchants
func (s *MerchantService) GetAllMerchants() ([]models.Merchant, error) {
	return s.repository.GetAllMerchants()
}

// obtiene un merchant por ID
func (s *MerchantService) GetMerchantByID(id uuid.UUID) (*models.Merchant, error) {
	return s.repository.GetMerchantByID(id)
}

// actualiza un merchant existente por su ID
func (s *MerchantService) UpdateMerchant(id uuid.UUID, merchant *models.Merchant) error {
	// busca si merchant existe
	existingMerchant, err := s.repository.GetMerchantByID(id)
	if err != nil {
		return fmt.Errorf("comercio no encontrado: %w", err)
	}

	// valida comision
	if merchant.Commission < 1.0 || merchant.Commission > 100.0 {
		return fmt.Errorf("la comisión debe estar entre 1.0 y 100.0")
	}
	// update merchant
	existingMerchant.Name = merchant.Name
	existingMerchant.Commission = merchant.Commission

	return s.repository.UpdateMerchant(id, existingMerchant)
}
