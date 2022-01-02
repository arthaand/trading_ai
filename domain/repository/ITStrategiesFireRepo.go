package repository

import (
	"trading-ai/domain/repository/entity"

	"gorm.io/gorm"
)

type ITStrategiesFireRepo interface {
	FindOne(entity.TStrategiesFire, string) (*entity.TStrategiesFire, error)
	FindOneWithPreload(entity.TStrategiesFire) (*entity.TStrategiesFire, error)
	FindAll(entity.TStrategiesFire, string, int, bool) (*[]entity.TStrategiesFire, error)
	Save(*entity.TStrategiesFire) error
	Update(*entity.TStrategiesFire, ...map[string]string) error
	SaveTx(*gorm.DB, *entity.TStrategiesFire) error
}
