package repositories

import (
	"myPagosApp/internal/models"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type MerchantRepository struct {
    DB *gorm.DB
}

// NewMerchantRepository crea un nuevo repositorio de comercios
func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
    return &MerchantRepository{DB: db}
}

// Crea nuevo merchant
func (r *MerchantRepository) CreateMerchant(merchant *models.Merchant) error {
    return r.DB.Create(merchant).Error
}

// GetAllMerchants obtiene todos los comercios de la base de datos
func (r *MerchantRepository) GetAllMerchants() ([]models.Merchant, error) {
    var merchants []models.Merchant
    err := r.DB.Find(&merchants).Error
    return merchants, err
}

// Actualizar un comercio por su ID
func (r *MerchantRepository) UpdateMerchant(id uuid.UUID, updatedMerchant *models.Merchant) error {
	return r.DB.Model(&models.Merchant{}).Where("id = ?", id).Updates(updatedMerchant).Error
}

// GetMerchantByID busca un comercio por su UUID
func (r *MerchantRepository) GetMerchantByID(id uuid.UUID) (*models.Merchant, error) {
    var merchant models.Merchant
    if err := r.DB.First(&merchant, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &merchant, nil
}


// UpdateMerchant actualiza un comercio existente en la base de datos
// func (r *MerchantRepository) UpdateMerchant(merchant *models.Merchant) error {
//     return r.DB.Save(merchant).Error
// }