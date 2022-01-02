package repository

import (
	"trading-ai/domain/repository/entity"

	"gorm.io/gorm"
)

type ITCurrencyHistoriesRepo interface {
	FindOne(entity.TCurrencyHistories, string) (*entity.TCurrencyHistories, error)
	FindOneWithPreload(entity.TCurrencyHistories) (*entity.TCurrencyHistories, error)
	FindAll(entity.TCurrencyHistories, string, int) (*[]entity.TCurrencyHistories, error)
	Save(*entity.TCurrencyHistories) error
	Update(*entity.TCurrencyHistories) error
	SaveTx(*gorm.DB, *entity.TCurrencyHistories) error
}
