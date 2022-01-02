package forex

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"trading-ai/constants"
	"trading-ai/domain/model/fastforex"
	"trading-ai/domain/repository"
	"trading-ai/domain/repository/entity"

	"andromeda.ottopay.id/ottopoint/ottopoint-shared/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type tradingForexUsecase struct {
	currencyHistories repository.ITCurrencyHistoriesRepo
	strategiesFire    repository.ITStrategiesFireRepo
	strategies        repository.IMStrategiesRepo
}

type ITradingForexUsecase interface {
	ReqTradingForex(*fiber.Ctx) *entity.TCurrencyHistories
}

func NewTradingForexUsecase(currencyHistories repository.ITCurrencyHistoriesRepo, strategiesFire repository.ITStrategiesFireRepo, strategies repository.IMStrategiesRepo) ITradingForexUsecase {
	return &tradingForexUsecase{currencyHistories: currencyHistories, strategiesFire: strategiesFire, strategies: strategies}
}

func (u *tradingForexUsecase) ReqTradingForex(c *fiber.Ctx) *entity.TCurrencyHistories {

	currHistory := Fastforex(c)

	//check if price same
	HistoryCheck := new(entity.TCurrencyHistories)
	HistoryCheck.Price = currHistory.Price
	datas, err := u.currencyHistories.FindOne(*HistoryCheck, "price_updated desc")

	if err != nil {
		//save new data histories
		fmt.Println("Error", err.Error())

		currHistory.Id = uuid.New()
		currHistory.CreatedAt = time.Now()

		fmt.Println("PairCurr", currHistory.MCurrencyCode)
		fmt.Println("Price", currHistory.Price)
		fmt.Println("history", currHistory)

		err = u.currencyHistories.Save(&currHistory)

		if err != nil {
			fmt.Printf(err.Error())
		}
	} else {
		fmt.Println("Data Found", datas.Id)
	}

	strategiesFire(u, currHistory)

	return &currHistory
}

func Fastforex(c *fiber.Ctx) entity.TCurrencyHistories {
	client := fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		DisablePathNormalizing:   true,
	}
	// init request object from fasthttp
	req := c.Request()
	req.SetRequestURI("https://api.fastforex.io/fetch-one?from=EUR&to=USD&api_key=55fb028cfa-76a04ac9ea-r4scea")

	// init response object from fasthttp
	res := c.Response()
	if err := client.Do(req, res); err != nil {
		fmt.Printf("Error")
	}

	HistoryFetch := fastforex.History{}
	currHistory := new(entity.TCurrencyHistories)
	// get responseCode
	err := json.Unmarshal(res.Body(), &HistoryFetch)
	if err != nil {
		fmt.Println(err)

	}

	Result := HistoryFetch.Result.(map[string]interface{})
	var PairCurr string
	var Price float64
	for key, value := range Result {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(float64))
		PairCurr = HistoryFetch.Base + key
		Price = value.(float64)
	}

	priceUpdated, _ := time.ParseInLocation(constant.DATE_TIME_FORMAT, HistoryFetch.Updated, time.Local)
	priceUpdated = priceUpdated.Add(time.Hour * 7)

	//set response to tempHistory
	currHistory = new(entity.TCurrencyHistories)
	currHistory.MCurrencyCode = PairCurr
	currHistory.Price = Price
	currHistory.PriceUpdated = priceUpdated

	return *currHistory
}

func strategiesFire(u *tradingForexUsecase, currHistory entity.TCurrencyHistories) {
	// Build Strategies
	HistoryCheck2 := new(entity.TCurrencyHistories)
	HistoryCheck2.MCurrencyCode = currHistory.MCurrencyCode
	HistoryCheck2.Status = 1
	datas2, _ := u.currencyHistories.FindAll(*HistoryCheck2, "price_updated desc", 4)
	fmt.Println("datas2 FindAll", datas2)

	var CountUp float64 = 0
	var CountDown float64 = 0
	var FirstVal float64
	var LastVal float64

	for i, r := range *datas2 {
		fmt.Println(i)
		if i == 0 {
			FirstVal = r.Price
			LastVal = r.Price
		} else {
			if FirstVal > r.Price {
				CountUp++
				fmt.Println("CountUp", CountUp)
			} else {
				fmt.Println("FirstVal", FirstVal, "BeforeVal", r.Price)
				CountDown++
			}
			FirstVal = r.Price
		}
		r.Status = 2
		updateCurrHistory(u, r)
	}
	fmt.Println(FirstVal)
	fmt.Println("CountUp", CountUp)
	fmt.Println("CountDown", CountDown)

	var Tren string
	var Action string
	var ConfidenceLevel float64
	var TpPercentage = 0.0003
	var ClPercentage = 0.0001
	var PriceTP float64
	var PriceCL float64
	if CountUp > CountDown {
		Tren = constants.Uptrend
		ConfidenceLevel = CountUp - CountDown
		fmt.Println("ConfidenceLevel a", ConfidenceLevel)
		ConfidenceLevel = (ConfidenceLevel / CountUp) * 100
		fmt.Println("ConfidenceLevel b", ConfidenceLevel)
		Action = constants.Buy
		// calculate TP CL
		PriceTP = LastVal * TpPercentage
		fmt.Println("PriceTP a", PriceTP)
		PriceTP = LastVal + PriceTP
		fmt.Println("PriceTP b", PriceTP)

		PriceCL = LastVal * ClPercentage
		PriceCL = LastVal - PriceCL
	} else {
		Tren = constants.Downtrend
		ConfidenceLevel = CountDown - CountUp
		ConfidenceLevel = (ConfidenceLevel / CountDown) * 100
		Action = constants.Sell

		PriceTP = LastVal * TpPercentage
		PriceTP = LastVal - PriceTP

		PriceCL = LastVal * ClPercentage
		PriceCL = LastVal + PriceCL
	}

	fmt.Println("Tren", Tren, "ConfidenceLevel", ConfidenceLevel)
	if ConfidenceLevel > float64(constants.ConfidenceLevel) {
		strategiesFire := new(entity.TStrategiesFire)
		strategiesFire.Id = uuid.New()
		strategiesFire.CreatedAt = time.Now()
		strategiesFire.Action = Action
		strategiesFire.MCurrencyCode = currHistory.MCurrencyCode
		strategiesFire.Price = LastVal
		strategiesFire.MStrategiesCode = "ST001"
		strategiesFire.Trend = Tren
		PriceTP, _ := strconv.ParseFloat(fmt.Sprintf("%5.5f", PriceTP), 8)
		strategiesFire.Tp = PriceTP
		PriceCL, _ := strconv.ParseFloat(fmt.Sprintf("%5.5f", PriceCL), 8)
		strategiesFire.Cl = PriceCL
		strategiesFire.ConfidenceLevel = ConfidenceLevel

		u.strategiesFire.Save(strategiesFire)

		for _, r := range *datas2 {
			//update MStrategiesFireId
			r.MStrategiesFireId = strategiesFire.Id
			updateCurrHistory(u, r)
		}
	}

}

func updateCurrHistory(u *tradingForexUsecase, currHistory entity.TCurrencyHistories) {
	currHistory.Status = 2
	u.currencyHistories.Update(&currHistory)
}
