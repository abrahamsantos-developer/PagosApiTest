package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"myPagosApp/internal/models"
)

type MerchantRepository struct {
	DB *gorm.DB
}

// crea un nuevo repositorio de merchants
func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{DB: db}
}

// crea nuevo merchant
func (r *MerchantRepository) CreateMerchant(merchant *models.Merchant) error {
	return r.DB.Create(merchant).Error
}

// obtiene todos los merchants DB
func (r *MerchantRepository) GetAllMerchants() ([]models.Merchant, error) {
	var merchants []models.Merchant
	err := r.DB.Find(&merchants).Error
	return merchants, err
}

// actualiza merchant por ID
func (r *MerchantRepository) UpdateMerchant(id uuid.UUID, updatedMerchant *models.Merchant) error {
	return r.DB.Model(&models.Merchant{}).Where("id = ?", id).Updates(updatedMerchant).Error
}

// obtiene merchant por ID
func (r *MerchantRepository) GetMerchantByID(id uuid.UUID) (*models.Merchant, error) {
	var merchant models.Merchant
	if err := r.DB.First(&merchant, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}
