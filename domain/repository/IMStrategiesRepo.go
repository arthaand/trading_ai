package repository

import (
	"trading-ai/domain/repository/entity"

	"gorm.io/gorm"
)

type IMStrategiesRepo interface {
	FindOne(entity.MStrategies, string) (*entity.MStrategies, error)
	FindOneWithPreload(entity.MStrategies) (*entity.MStrategies, error)
	FindAll(entity.MStrategies, string, int, bool) (*[]entity.MStrategies, error)
	Save(*entity.MStrategies) error
	Update(*entity.MStrategies, ...map[string]string) error
	SaveTx(*gorm.DB, *entity.MStrategies) error
}
