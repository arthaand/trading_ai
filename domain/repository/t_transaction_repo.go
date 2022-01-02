package repository

import (
	"andromeda.ottopay.id/ottopoint/ottopoint-orm/entity"
	"gorm.io/gorm"
)

type ITTransactionRepo interface {
	FindOne(entity.TTransaction) (*entity.TTransaction, error)
	FindAll(entity.TTransaction) (*[]entity.TTransaction, error)
	FindAllPagination(limit, page int64, qry ...entity.TTransaction) (*[]entity.TTransaction, int64, error)
	SaveTx(*entity.TTransaction) (*gorm.DB, error)
}
