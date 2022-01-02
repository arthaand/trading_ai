package persistence

import (
	"errors"
	"trading-ai/domain/repository"
	"trading-ai/domain/repository/entity"

	"gorm.io/gorm"
)

type MStrategiesRepo struct {
	db *gorm.DB
}

func newMStrategiesRepo(db *gorm.DB) repository.IMStrategiesRepo {
	return &MStrategiesRepo{db: db}
}

func (r *MStrategiesRepo) FindAll(qry entity.MStrategies, orderBy string, limit int, iscall bool) (*[]entity.MStrategies, error) {
	data := new([]entity.MStrategies)

	if err := r.db.Order(orderBy).Where(qry).Limit(limit).Find(&data).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *MStrategiesRepo) FindOneWithPreload(qry entity.MStrategies) (*entity.MStrategies, error) {
	data := new(entity.MStrategies)
	if errors.Is(r.db.Preload("Transaction").Where(qry).First(&data).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *MStrategiesRepo) FindOne(qry entity.MStrategies, orderBy string) (*entity.MStrategies, error) {
	data := new(entity.MStrategies)
	if errors.Is(r.db.Order(orderBy).Where(qry).First(&data).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *MStrategiesRepo) Save(data *entity.MStrategies) error {
	if err := r.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

// Update update by primary key with struct and optional query map
func (r *MStrategiesRepo) Update(data *entity.MStrategies, qry ...map[string]string) error {
	db := r.db
	if len(qry) > 0 {
		db = r.db.Where(qry[0])
	}
	if err := db.Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *MStrategiesRepo) SaveTx(tx *gorm.DB, data *entity.MStrategies) error {
	if err := tx.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
