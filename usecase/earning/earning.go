package earning

// import (
// 	"trading-ai/domain/repository"
// 	"time"

// 	kafka "andromeda.ottopay.id/ottopoint/ottopoint-kafka"
// 	"andromeda.ottopay.id/ottopoint/ottopoint-orm/entity"
// 	"andromeda.ottopay.id/ottopoint/ottopoint-shared/constant"
// 	"andromeda.ottopay.id/ottopoint/ottopoint-shared/model"
// 	"github.com/Shopify/sarama"
// 	"github.com/google/uuid"
// 	"github.com/jinzhu/copier"
// )

// type earningUsecase struct {
// 	earning     repository.ITEarningRepo
// 	transaction repository.ITTransactionRepo
// 	producer    sarama.AsyncProducer
// }

// type IEarningUsecase interface {
// 	ReqEarning(*model.ReqEarningHandler, string) (*model.BaseResp, *entity.TEarning)
// 	PublishReqEarning(*entity.TEarning, string)
// }

// func NewEarningUsecase(
// 	earn repository.ITEarningRepo,
// 	trx repository.ITTransactionRepo,
// 	producer sarama.AsyncProducer,
// ) IEarningUsecase {
// 	return &earningUsecase{earning: earn, transaction: trx,
// 		producer: producer}
// }

// // ReqEarning request from apps loyalty to ottostamp -> trading-ai
// func (u *earningUsecase) ReqEarning(req *model.ReqEarningHandler, institutionId string) (*model.BaseResp, *entity.TEarning) {
// 	res := new(model.BaseResp).OK()
// 	// check if QR is exists
// 	earnId, _ := uuid.Parse(req.EarningId)
// 	qry := entity.TEarning{}
// 	qry.Id = earnId
// 	data, err := u.earning.FindOne(qry)
// 	if err != nil {
// 		res.ResponseCode = constant.RC_QR_NOT_FOUND
// 		res.ResponseDesc = constant.RD_QR_NOT_FOUND
// 		return res, nil
// 	}
// 	// check if qr is not yet redeemed
// 	if data.StatusCd != constant.RC_PENDING {
// 		res.ResponseCode = constant.RC_QR_HAS_BEEN_REDEEMED
// 		res.ResponseDesc = constant.RD_QR_HAS_BEEN_REDEEMED
// 		return res, nil
// 	}
// 	// update t_earning
// 	reqAt, _ := time.Parse(constant.DATE_TIME_FORMAT, req.RequestAt)
// 	data.RequestorInstitutionId = institutionId
// 	data.RequestAt = &reqAt
// 	data.AccountNo = req.AccountNo
// 	if err = u.earning.Update(data); err != nil {
// 		return res.Err(err), nil
// 	}

// 	return new(model.BaseResp).OK(), data
// }

// func (u *earningUsecase) PublishReqEarning(data *entity.TEarning, reqTopic string) {
// 	// init response
// 	resPayload := model.EarnPubPayload{}
// 	// get trx and trx detail data based on invoice no
// 	qryTrx := new(entity.TTransaction)
// 	qryTrx.Id = data.TTransactionId
// 	trx, _ := u.transaction.FindOne(*qryTrx)

// 	// dto
// 	copier.Copy(&resPayload, data)
// 	copier.Copy(&resPayload, trx)
// 	if data.EarningCode == "" {
// 		copier.Copy(&resPayload.DetailProduct, trx.Detail)
// 	}

// 	// publish
// 	kafka.PublishAsync(u.producer, resPayload, reqTopic)
// }
