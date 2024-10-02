package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"myPagosApp/internal/models"
)

type TransactionRepository struct {
	DB *gorm.DB
}

// crea nuevo repositorio de transactions
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// crea nueva transaction
func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

// obtiene todas las transactions de DB
func (r *TransactionRepository) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Find(&transactions).Error
	return transactions, err
}

// obtiene todas las transactions de un merchant por ID
func (r *TransactionRepository) GetTransactionsByMerchantID(merchantID uuid.UUID) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Where("merchant_id = ?", merchantID).Find(&transactions).Error
	return transactions, err
}

// obtiene una transaction por ID de la misma
func (r *TransactionRepository) GetTransactionByID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.First(&transaction, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

// suma todas las comisiones de todas las transactions
func (r *TransactionRepository) GetTotalProfits() (float64, error) {
	var totalProfits float64
	// Suma todos los fees de las transacciones
	err := r.DB.Model(&models.Transaction{}).Select("SUM(fee)").Scan(&totalProfits).Error
	if err != nil {
		return 0, err
	}
	return totalProfits, nil
}

// suma todas las fees de las transactions por merchantID 
func (r *TransactionRepository) SumCommissionsByMerchantID(merchantID uuid.UUID) (float64, error) {
	var totalCommission float64
	err := r.DB.Model(&models.Transaction{}).
		Where("merchant_id = ?", merchantID).
		Select("SUM(commission)").
		Scan(&totalCommission).Error
	if err != nil {
		return 0, err
	}
	return totalCommission, nil
}
