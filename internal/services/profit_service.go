package services

import (
	"myPagosApp/internal/repositories"
	"github.com/google/uuid"
)

// ProfitService define la lógica para calcular las ganancias
type ProfitService struct {
	transactionRepo *repositories.TransactionRepository
}

// NewProfitService crea un nuevo servicio de ganancias
func NewProfitService(transactionRepo *repositories.TransactionRepository) *ProfitService {
	return &ProfitService{transactionRepo: transactionRepo}
}

// GetTotalProfits calcula las ganancias totales de todas las transacciones
func (s *ProfitService) GetTotalProfits() (float64, error) {
	return s.transactionRepo.GetTotalProfits()
}

// GetProfitsByMerchantID calcula las ganancias de un comercio específico
func (s *ProfitService) GetProfitsByMerchantID(merchantID uuid.UUID) (float64, error) {
	return s.transactionRepo.SumCommissionsByMerchantID(merchantID)
}
