package services

import (
	"fmt"
	"github.com/google/uuid"
	"myPagosApp/internal/models"
	"myPagosApp/internal/repositories"
)

// logica para las transactions
type TransactionService struct {
	repository   *repositories.TransactionRepository
	merchantRepo *repositories.MerchantRepository
}

// crea nuevo service de transactions
func NewTransactionService(repo *repositories.TransactionRepository, merchantRepo *repositories.MerchantRepository) *TransactionService {
	return &TransactionService{
		repository:   repo,
		merchantRepo: merchantRepo,
	}
}

// crea una nueva transaction. calcula el fee basado en la comision del merchant
func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	// obtiene merchant asociado por MerchantID
	merchant, err := s.merchantRepo.GetMerchantByID(transaction.MerchantID)
	if err != nil {
		return fmt.Errorf("comercio no encontrado: %w", err)
	}

	// calcula la comission y fee
	transaction.Commission = merchant.Commission
	transaction.Fee = (transaction.Amount * transaction.Commission) / 100.0

	// guarda transaccion
	return s.repository.CreateTransaction(transaction)
}

// obtiene todas las transactions
func (s *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repository.GetAllTransactions()
}

// obtiene todas las transacctions by MerchantId
func (s *TransactionService) GetTransactionsByMerchantID(merchantID uuid.UUID) ([]models.Transaction, error) {
	return s.repository.GetTransactionsByMerchantID(merchantID)
}

// obtiene una transaction por su ID
func (s *TransactionService) GetTransactionByID(id uuid.UUID) (*models.Transaction, error) {
	return s.repository.GetTransactionByID(id)
}
