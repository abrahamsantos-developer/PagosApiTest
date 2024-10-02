package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// representa un merchant
type Merchant struct {
	ID           uuid.UUID     `gorm:"type:uuid;primary_key" json:"id"`
	Name         string        `gorm:"size:255" json:"name"`
	Commission   float64       `json:"commission"`
	Transactions []Transaction `gorm:"foreignKey:MerchantID" json:"transactions"` //agregamos la relacion
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

// genera UUID antes de insertar nuevo registro
func (m *Merchant) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
