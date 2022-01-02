package persistence

import (
	"errors"
	"trading-ai/domain/repository"
	"trading-ai/domain/repository/entity"

	"gorm.io/gorm"
)

type TStrategiesFireRepo struct {
	db *gorm.DB
}

func newTStrategiesFireRepo(db *gorm.DB) repository.ITStrategiesFireRepo {
	return &TStrategiesFireRepo{db: db}
}

func (r *TStrategiesFireRepo) FindAll(qry entity.TStrategiesFire, orderBy string, limit int, iscall bool) (*[]entity.TStrategiesFire, error) {
	data := new([]entity.TStrategiesFire)

	if err := r.db.Order(orderBy).Where(qry).Limit(limit).Find(&data).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *TStrategiesFireRepo) FindOneWithPreload(qry entity.TStrategiesFire) (*entity.TStrategiesFire, error) {
	data := new(entity.TStrategiesFire)
	if errors.Is(r.db.Preload("Transaction").Where(qry).First(&data).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *TStrategiesFireRepo) FindOne(qry entity.TStrategiesFire, orderBy string) (*entity.TStrategiesFire, error) {
	data := new(entity.TStrategiesFire)
	if errors.Is(r.db.Order(orderBy).Where(qry).First(&data).Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (r *TStrategiesFireRepo) Save(data *entity.TStrategiesFire) error {
	if err := r.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

// Update update by primary key with struct and optional query map
func (r *TStrategiesFireRepo) Update(data *entity.TStrategiesFire, qry ...map[string]string) error {
	db := r.db
	if len(qry) > 0 {
		db = r.db.Where(qry[0])
	}
	if err := db.Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *TStrategiesFireRepo) SaveTx(tx *gorm.DB, data *entity.TStrategiesFire) error {
	if err := tx.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
