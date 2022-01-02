package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	docs "trading-ai/docs"
	"trading-ai/domain/model"
	"trading-ai/infrastructure/persistence"
	httpHandler "trading-ai/interfaces/http"
	forexUc "trading-ai/usecase/forex"

	shared "andromeda.ottopay.id/ottopoint/ottopoint-shared"
	constant "andromeda.ottopay.id/ottopoint/ottopoint-shared/constant"
	m "andromeda.ottopay.id/ottopoint/ottopoint-shared/model"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/astaxie/beego/toolbox"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/jinzhu/configor"
)

// @contact.name Bumi Langit Team
// @contact.url https://ottopoint.id
// @contact.email info@ottopoint.id
// @license.name Trading AI License
func main() {
	// init config using yml
	err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: false}).Load(&shared.Config, "config.yml")
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	// init repository
	repo, err := persistence.NewRepositories()
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	// init usecase
	// usecase using producer async for centralized logging (plan) in cloud native way
	//qrstringUc := qrUsecase.NewQrstringUsecase(repo.Earning, repo.Transaction, repo.TransactionDt)
	forexUc := forexUc.NewTradingForexUsecase(repo.CurrencyHistories, repo.StrategiesFire, repo.Strategies)

	// earningUc := earningUsecase.NewEarningUsecase(repo.Earning, repo.Transaction, infr

	// init http handler
	//posHandler := httpHandler.NewQrstringHandler(qrstringUc)
	forexHandler := httpHandler.NewTradingForexHandler(forexUc)

	// init kafka handler
	//earningKafkaHandler := kafkaHandler.NewEarningHandler(infrastructure.GetConsumer(), infrastructure.GetProducer(), statUpdUc)

	// init services with kafka
	wg := sync.WaitGroup{}
	wg.Add(1)
	//earningKafkaHandler.EarningStatusUpdate(topic.POS_REQ_EARNING_CREATED, topic.POS_REQ_EARNING_UPDATED)
	// wg.Wait() without wait because app.Listen is also waitgroup

	// init services with http (gofiber)
	// init server with config
	cfgFiber := fiber.Config{
		BodyLimit: 5 * 1024 * 1024, // 100 mb limit upload
	}
	app := fiber.New(cfgFiber)

	// init middleware
	// logger
	if shared.Config.Env == constant.ENV_DEV {
		app.Use(logger.New())
	}
	// recover
	app.Use(recover.New())
	// cors default
	app.Use(cors.New())
	// helmet default
	app.Use(helmet.New())
	// init swagger
	app.Use("/swagger", swagger.Handler)
	// init swagger header
	docs.SwaggerInfo.Title = "OttoPoint POS Service Swagger UI"
	docs.SwaggerInfo.Version = shared.Config.Version
	docs.SwaggerInfo.Description = "API Documentation Spesification at OttoPoint POS Service using Swagger UI version 2"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = shared.Config.Host
	if shared.Config.Port != "" {
		docs.SwaggerInfo.Host = shared.Config.Host + ":" + shared.Config.Port
	}
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// init rest handler
	// POS SERVICE
	pos := app.Group("trading-ai")
	// add health check to trading-ai
	pos.Get("/v1/healthcheck", func(c *fiber.Ctx) error {
		return model.Response(c, &m.BaseResp{
			ResponseCode: constant.RC_SUCCESS,
			ResponseDesc: constant.RD_SUCCESS,
			Code:         http.StatusOK,
		})
	})
	//pos.Post("/v1/earning/qrstring", middleware.SignatureAuth, posHandler.ReqQrString)
	pos.Get("/v1/earning/status", forexHandler.ReqTradingForex)

	// run server
	app.Listen(":" + shared.Config.Port)
}

func init() {
	// schedulers := toolbox.NewTask("schedulers", "0 0/5 * * * *", func() error {
	// 	// schedulers := toolbox.NewTask("schedulers", fmt.Sprintf("%v", timeSchedule), func() error {

	// 	ReqTradingForex()
	// 	return nil
	// })

	perMinuteTask := toolbox.NewTask("perMinuteTask", "0 */1 * * * *", func() error {
		fmt.Printf("time: %v it is per minute task\n", time.Now().Format("2006-01-02 15:04:05"))
		http.Get("http://localhost:8082/trading-ai/v1/earning/status")
		return nil
	})

	//schedulers.Run()

	toolbox.AddTask("perMinuteTask", perMinuteTask)
	toolbox.StartTask()

}
