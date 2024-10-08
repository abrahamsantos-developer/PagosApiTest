package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// representa transaction de un merchant
type Transaction struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	MerchantID uuid.UUID `gorm:"type:uuid" json:"merchant_id"`
	Amount     float64   `json:"amount"`
	Commission float64   `json:"commission"` // porcentaje aplicado (tomado del comercio)
	Fee        float64   `json:"fee"`        // comisión calculada
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// genera UUID antes de crear transaction
func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
