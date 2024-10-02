package services

import (
	"github.com/google/uuid"
	"myPagosApp/internal/repositories"
)

// logica para calcular las profits
type ProfitService struct {
	transactionRepo *repositories.TransactionRepository
}

// crea un nuevo service de profits
func NewProfitService(transactionRepo *repositories.TransactionRepository) *ProfitService {
	return &ProfitService{transactionRepo: transactionRepo}
}

// calcula las profits totales de todas las transactions
func (s *ProfitService) GetTotalProfits() (float64, error) {
	return s.transactionRepo.GetTotalProfits()
}

// calcula las ganancias de un merchant por su ID
func (s *ProfitService) GetProfitsByMerchantID(merchantID uuid.UUID) (float64, error) {
	return s.transactionRepo.SumCommissionsByMerchantID(merchantID)
}
