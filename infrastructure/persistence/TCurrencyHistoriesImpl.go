package persistence

import (
	"errors"
	"fmt"
	"trading-ai/domain/repository"
	"trading-ai/domain/repository/entity"

	"gorm.io/gorm"
)

type MCurrencyHistoriesRepo struct {
	db *gorm.DB
}

func newCurrencyHistoriesRepo(db *gorm.DB) repository.ITCurrencyHistoriesRepo {
	return &MCurrencyHistoriesRepo{db: db}
}

func (r *MCurrencyHistoriesRepo) FindAll(qry entity.TCurrencyHistories, orderBy string, limit int) (*[]entity.TCurrencyHistories, error) {
	data := new([]entity.TCurrencyHistories)
	fmt.Printf("aaaa", qry)
	if err := r.db.Order(orderBy).Where(qry).Limit(limit).Find(&data).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *MCurrencyHistoriesRepo) FindOneWithPreload(qry entity.TCurrencyHistories) (*entity.TCurrencyHistories, error) {
	data := new(entity.TCurrencyHistories)
	if errors.Is(r.db.Preload("Transaction").Where(qry).First(&data).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *MCurrencyHistoriesRepo) FindOne(qry entity.TCurrencyHistories, orderBy string) (*entity.TCurrencyHistories, error) {
	data := new(entity.TCurrencyHistories)
	if errors.Is(r.db.Order(orderBy).Where(qry).First(&data).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *MCurrencyHistoriesRepo) Save(data *entity.TCurrencyHistories) error {
	if err := r.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

// Update update by primary key with struct and optional query map
func (r *MCurrencyHistoriesRepo) Update(data *entity.TCurrencyHistories) error {
	db := r.db
	// if len(qry) > 0 {
	// 	db = r.db.Where(qry[0])
	// }
	if err := db.Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *MCurrencyHistoriesRepo) SaveTx(tx *gorm.DB, data *entity.TCurrencyHistories) error {
	if err := tx.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
