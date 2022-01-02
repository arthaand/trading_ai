package http

import (
	"net/http"

	usecase "trading-ai/usecase/forex"

	"github.com/gofiber/fiber/v2"
)

type tradingForexHandler struct {
	tradingForex usecase.ITradingForexUsecase
}

func NewTradingForexHandler(trading usecase.ITradingForexUsecase) *tradingForexHandler {
	return &tradingForexHandler{tradingForex: trading}
}

// CheckStatus godoc
// @Summary Check Status Hit by ottopoint-stamp
// @Id CheckStatus
// @Tags Pos Service
// @Accept json
// @Produce json
// @Param Institutionid header string true "UUID Partner (example = 'PSM0002')"
// @Param Deviceid header string true "Hardcoded"
// @Param Geolocation header string true " "
// @Param Channelid header string true "Hardcoded (example = 'H2H')"
// @Param Signature header string true "Generated system by partner, all request POST validate with signature (example = 'a0acd4cc5bb2')"
// @Param Appsid header string true "(Example = 'pos-customer')"
// @Param Authorization header string true " "
// @Param Timestamp header string true "(Example = '1579666534')"
// @Param payload body model.ReqCheckStatus true " "
// @Success 200 {object} model.BaseResp " "
// @Router /trading-ai/v1/earning/status [post]
func (h *tradingForexHandler) ReqTradingForex(c *fiber.Ctx) error {

	res := h.tradingForex.ReqTradingForex(c)

	return c.Status(http.StatusOK).JSON(res)
	//return res
}
