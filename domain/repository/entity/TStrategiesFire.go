package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TStrategiesFire struct {
	Id              uuid.UUID `gorm:"primaryKey;type:uuid"` // UUID
	MCurrencyCode   string
	MStrategiesCode string
	Action          string
	Trend           string
	Price           float64
	Tp              float64
	Cl              float64
	Result          string
	WinLosePip      float64
	ConfidenceLevel float64
	CreatedAt       time.Time
}

// BeforeCreate hook for uuid
func (te *TStrategiesFire) BeforeCreate(tx *gorm.DB) (err error) {
	te.Id = uuid.New()
	return
}
