package models

import (
	"time"
	
)


// Merchant representa la estructura de un comercio
type Merchant struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"size:255" json:"name"`
    Commission uint      `json:"commission"` //  porcentaje comision
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
