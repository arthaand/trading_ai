package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TCurrencyHistories struct {
	Id                uuid.UUID `gorm:"primaryKey;type:uuid"` // UUID
	MCurrencyCode     string
	Price             float64
	CreatedAt         time.Time
	PriceUpdated      time.Time
	Status            int8
	UpdatedAt         time.Time
	MStrategiesFireId uuid.UUID
}

// BeforeCreate hook for uuid
func (te *TCurrencyHistories) BeforeCreate(tx *gorm.DB) (err error) {
	te.Id = uuid.New()
	return
}
