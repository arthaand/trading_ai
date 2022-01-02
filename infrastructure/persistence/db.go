package persistence

import (
	"log"
	"trading-ai/domain/repository"

	orm "andromeda.ottopay.id/ottopoint/ottopoint-orm"
	shared "andromeda.ottopay.id/ottopoint/ottopoint-shared"
	"gorm.io/gorm"
)

type repositories struct {
	db *gorm.DB
	//CandleCurrencyHistories repository.IMCandleCurrencyHistoriesRepo
	CurrencyHistories repository.ITCurrencyHistoriesRepo
	StrategiesFire    repository.ITStrategiesFireRepo
	Strategies        repository.IMStrategiesRepo
	// Transaction             repository.ITTransactionRepo
	// TransactionDt           repository.ITTransactionDtRepo
}

func NewRepositories() (*repositories, error) {
	// init connection postgres
	db, err := orm.InitPostgres(shared.Config.Env,
		shared.Config.Postgres.Host, shared.Config.Postgres.User,
		shared.Config.Postgres.Password, shared.Config.Postgres.Name,
		shared.Config.Postgres.Port)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	log.Printf("INFO: Connected to DB")
	return &repositories{
		db: db,
		//CurrencyHistories:     newMCurrencyHistoriesRepoRepo(db),
		CurrencyHistories: newCurrencyHistoriesRepo(db),
		StrategiesFire:    newTStrategiesFireRepo(db),
		Strategies:        newMStrategiesRepo(db),
		// Transaction:       newTransactionRepo(db),
		// TransactionDt:     newTransactionDt(db),
	}, nil
}
