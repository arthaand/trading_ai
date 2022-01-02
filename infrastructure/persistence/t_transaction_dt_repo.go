package persistence

import (
	"errors"
	"trading-ai/domain/repository"

	"andromeda.ottopay.id/ottopoint/ottopoint-orm/entity"
	"gorm.io/gorm"
)

type tTransactionDtRepo struct {
	db *gorm.DB
}

func newTransactionDt(db *gorm.DB) repository.ITTransactionDtRepo {
	return &tTransactionDtRepo{db: db}
}

func (r *tTransactionDtRepo) FindOne(qry entity.TTransactionDt) (*entity.TTransactionDt, error) {
	data := new(entity.TTransactionDt)
	if errors.Is(r.db.Where(qry).Find(&data).Error, gorm.ErrRecordNotFound) {
		return data, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *tTransactionDtRepo) SaveTx(tx *gorm.DB, data *entity.TTransactionDt) error {
	if err := tx.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
