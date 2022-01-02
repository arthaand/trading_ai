package http

// import (
// 	"net/http"
// 	m "trading-ai/domain/model"
// 	usecase "trading-ai/usecase/earning"

// 	"andromeda.ottopay.id/ottopoint/ottopoint-kafka/topic"
// 	"andromeda.ottopay.id/ottopoint/ottopoint-shared/common"
// 	"andromeda.ottopay.id/ottopoint/ottopoint-shared/model"
// 	"github.com/gofiber/fiber/v2"
// )

// type earningHandler struct {
// 	earning usecase.IEarningUsecase
// }

// func NewEarningHandler(earning usecase.IEarningUsecase) *earningHandler {
// 	return &earningHandler{earning: earning}
// }

// // ReqEarning godoc
// // @Summary hit from apps mobile / partner user apps
// // @Id ReqEarning
// // @Tags Pos Service
// // @Accept json
// // @Produce json
// // @Param Institutionid header string true "UUID Partner (example = 'PSM0002')"
// // @Param Deviceid header string true "Hardcoded"
// // @Param Geolocation header string true " "
// // @Param Channelid header string true "Hardcoded (example = 'H2H')"
// // @Param Signature header string true "Generated system by partner, all request POST validate with signature (example = 'a0acd4cc5bb2')"
// // @Param Appsid header string true "(Example = 'pos-customer')"
// // @Param Authorization header string true " "
// // @Param Timestamp header string true "(Example = '1579666534')"
// // @Param payload body model.ReqEarningHandler true "Request Body"
// // @Success 200 {object} model.BaseResp{data=model.ResQrString} " "
// // @Router /trading-ai/v1/earning/request [post]
// func (h *earningHandler) ReqEarning(c *fiber.Ctx) error {
// 	var reqHandler model.ReqEarningHandler
// 	c.BodyParser(&reqHandler)
// 	if err := common.Validate(reqHandler); err != nil {
// 		return m.ResponseErr(c, http.StatusBadRequest, err.Error())
// 	}
// 	resp, data := h.earning.ReqEarning(&reqHandler, c.Get("Institutionid"))
// 	// publish to kafka should use goroutine
// 	if data != nil {
// 		h.earning.PublishReqEarning(data, topic.POS_REQ_EARNING_REQUESTED)
// 		return m.Response(c, resp)
// 	}
// 	return m.ResponseErr(c, http.StatusUnprocessableEntity, resp.ResponseDesc)
// }
