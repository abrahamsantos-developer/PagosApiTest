package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Merchant representa la estructura de un comercio
type Merchant struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name       string    `gorm:"size:255" json:"name"`
	Commission uint      `json:"commission"` //  porcentaje comision
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//Genera UUID antes delnuevo registro
func (m *Merchant) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New() // Generar un nuevo UUID
	return
}
