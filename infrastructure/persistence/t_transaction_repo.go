package persistence

import (
	"errors"
	"trading-ai/domain/repository"

	orm "andromeda.ottopay.id/ottopoint/ottopoint-orm"
	"andromeda.ottopay.id/ottopoint/ottopoint-orm/entity"
	"gorm.io/gorm"
)

type tTransactionRepo struct {
	db *gorm.DB
}

func newTransactionRepo(db *gorm.DB) repository.ITTransactionRepo {
	return &tTransactionRepo{db: db}
}

func (r *tTransactionRepo) FindOne(qry entity.TTransaction) (*entity.TTransaction, error) {
	data := new(entity.TTransaction)
	if errors.Is(r.db.Preload("Detail").Where(qry).First(&data).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *tTransactionRepo) FindAll(qry entity.TTransaction) (*[]entity.TTransaction, error) {
	data := new([]entity.TTransaction)
	if err := r.db.Preload("Detail").Where(qry).Find(&data).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *tTransactionRepo) FindAllPagination(limit, page int64, qry ...entity.TTransaction) (*[]entity.TTransaction, int64, error) {
	data := new([]entity.TTransaction)
	stmt := r.db.Where(qry)
	if err := stmt.Preload("Detail").Scopes(orm.Paginate(page, limit)).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, stmt.RowsAffected, nil
}

func (r *tTransactionRepo) SaveTx(data *entity.TTransaction) (*gorm.DB, error) {
	tx := r.db.Begin()
	if err := tx.Create(&data).Error; err != nil {
		return nil, err
	}
	return tx, nil
}
