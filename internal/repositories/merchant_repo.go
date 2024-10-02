package repositories

import (
	"myPagosApp/internal/models"
	"gorm.io/gorm"
)

type MerchantRepository struct {
    DB *gorm.DB
}

// NewMerchantRepository crea un nuevo repositorio de comercios
func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
    return &MerchantRepository{DB: db}
}

// CreateMerchant guarda un nuevo comercio en la base de datos
func (r *MerchantRepository) CreateMerchant(merchant *models.Merchant) error {
    return r.DB.Create(merchant).Error
}

// GetAllMerchants obtiene todos los comercios de la base de datos
func (r *MerchantRepository) GetAllMerchants() ([]models.Merchant, error) {
    var merchants []models.Merchant
    err := r.DB.Find(&merchants).Error
    return merchants, err
}

// UpdateMerchant actualiza un comercio existente en la base de datos
func (r *MerchantRepository) UpdateMerchant(merchant *models.Merchant) error {
    return r.DB.Save(merchant).Error
}