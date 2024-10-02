package services

import (
	"fmt"
	"github.com/google/uuid"
	"myPagosApp/internal/models"
	"myPagosApp/internal/repositories"
)

// TransactionService define la lógica de negocio para las transacciones
type TransactionService struct {
	repository    *repositories.TransactionRepository
	merchantRepo  *repositories.MerchantRepository
}

// NewTransactionService crea un nuevo servicio de transacciones
func NewTransactionService(repo *repositories.TransactionRepository, merchantRepo *repositories.MerchantRepository) *TransactionService {
	return &TransactionService{
		repository:   repo,
		merchantRepo: merchantRepo,
	}
}

// CreateTransaction crea una nueva transacción, calculando el fee basado en la comisión del comercio
func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	// Obtener el comercio asociado por MerchantID
	merchant, err := s.merchantRepo.GetMerchantByID(transaction.MerchantID)
	if err != nil {
		return fmt.Errorf("comercio no encontrado: %w", err)
	}

	// Calcular la comisión y el fee
	transaction.Commission = merchant.Commission
	transaction.Fee = (transaction.Amount * transaction.Commission) / 100.0

	// Guardar la transacción
	return s.repository.CreateTransaction(transaction)
}

// GetAllTransactions obtiene todas las transacciones
func (s *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repository.GetAllTransactions()
}

// GetTransactionsByMerchantID obtiene todas las transacciones de un comercio específico
func (s *TransactionService) GetTransactionsByMerchantID(merchantID uuid.UUID) ([]models.Transaction, error) {
	return s.repository.GetTransactionsByMerchantID(merchantID)
}

// GetTransactionByID obtiene una transacción por su UUID
func (s *TransactionService) GetTransactionByID(id uuid.UUID) (*models.Transaction, error) {
	return s.repository.GetTransactionByID(id)
}
