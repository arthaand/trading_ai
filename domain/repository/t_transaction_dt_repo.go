package repository

import (
	"andromeda.ottopay.id/ottopoint/ottopoint-orm/entity"
	"gorm.io/gorm"
)

type ITTransactionDtRepo interface {
	FindOne(entity.TTransactionDt) (*entity.TTransactionDt, error)
	SaveTx(*gorm.DB, *entity.TTransactionDt) error
}
