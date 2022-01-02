package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MStrategies struct {
	Id              uuid.UUID `gorm:"primaryKey;type:uuid"` // UUID
	Code            string
	TpPercentage    int32
	ClPercentage    int32
	LookupHistories int32
	IsActive        bool
	IsReversal      string
	LookupMa        int32
	CreatedAt       time.Time
}

// BeforeCreate hook for uuid
func (te *MStrategies) BeforeCreate(tx *gorm.DB) (err error) {
	te.Id = uuid.New()
	return
}
