package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid"` // UUID
	CreatedAt time.Time
	// CreatedBy string
	// UpdatedAt *time.Time
	// UpdatedBy string
}

type MCandleCurrencyHistories struct {
	BaseEntity
	MCurrencyCode string
	open          float32
	low           float32
	high          float32
	close         float32
	openDate      time.Time
	closeDate     time.Time
}

// BeforeCreate hook for uuid
func (te *MCandleCurrencyHistories) BeforeCreate(tx *gorm.DB) (err error) {
	te.Id = uuid.New()
	return
}
