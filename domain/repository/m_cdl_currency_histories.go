package repository

import (
	"trading-ai/domain/repository/entity"

	"gorm.io/gorm"
)

type IMCandleCurrencyHistoriesRepo interface {
	FindOne(entity.MCandleCurrencyHistories) (*entity.MCandleCurrencyHistories, error)
	FindOneWithPreload(entity.MCandleCurrencyHistories) (*entity.MCandleCurrencyHistories, error)
	FindAll(entity.MCandleCurrencyHistories) (*[]entity.MCandleCurrencyHistories, error)
	Save(*entity.MCandleCurrencyHistories) error
	Update(*entity.MCandleCurrencyHistories, ...map[string]string) error
	SaveTx(*gorm.DB, *entity.MCandleCurrencyHistories) error
}
