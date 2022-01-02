package persistence

// import (
// 	"errors"
// 	"trading-ai/domain/repository"
// 	"trading-ai/domain/repository/entity"

// 	"gorm.io/gorm"
// )

// type mCurrencyHistoriesRepo struct {
// 	db *gorm.DB
// }

// func newMCurrencyHistoriesRepoRepo(db *gorm.DB) repository.IMCurrencyHistoriesRepo {
// 	return &mCurrencyHistoriesRepo{db: db}
// }

// func (r *mCurrencyHistoriesRepo) FindAll(qry entity.MCurrencyHistories) (*[]entity.MCurrencyHistories, error) {
// 	data := new([]entity.MCurrencyHistories)
// 	if err := r.db.Where(qry).Find(&data).Error; err != nil {
// 		return nil, gorm.ErrRecordNotFound
// 	}
// 	return data, nil
// }

// func (r *mCurrencyHistoriesRepo) FindOneWithPreload(qry entity.MCurrencyHistories) (*entity.MCurrencyHistories, error) {
// 	data := new(entity.MCurrencyHistories)
// 	if errors.Is(r.db.Preload("Transaction").Where(qry).First(&data).Error, gorm.ErrRecordNotFound) {
// 		return nil, gorm.ErrRecordNotFound
// 	}
// 	return data, nil
// }

// func (r *mCurrencyHistoriesRepo) FindOne(qry entity.MCurrencyHistories) (*entity.MCurrencyHistories, error) {
// 	data := new(entity.MCurrencyHistories)
// 	if errors.Is(r.db.Where(qry).First(&data).Error, gorm.ErrRecordNotFound) {
// 		return nil, gorm.ErrRecordNotFound
// 	}
// 	return data, nil
// }

// func (r *mCurrencyHistoriesRepo) Save(data *entity.MCurrencyHistories) error {
// 	if err := r.db.Create(&data).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Update update by primary key with struct and optional query map
// func (r *mCurrencyHistoriesRepo) Update(data *entity.MCurrencyHistories, qry ...map[string]string) error {
// 	db := r.db
// 	if len(qry) > 0 {
// 		db = r.db.Where(qry[0])
// 	}
// 	if err := db.Updates(&data).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *mCurrencyHistoriesRepo) SaveTx(tx *gorm.DB, data *entity.MCurrencyHistories) error {
// 	if err := tx.Create(&data).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
