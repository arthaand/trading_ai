package earning

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"trading-ai/domain/repository"

// 	kafka "andromeda.ottopay.id/ottopoint/ottopoint-kafka"
// 	"andromeda.ottopay.id/ottopoint/ottopoint-orm/entity"
// 	"andromeda.ottopay.id/ottopoint/ottopoint-shared/model"
// 	"github.com/jinzhu/copier"
// )

// type earningStatUpdUsecase struct {
// 	earning repository.ITEarningRepo
// 	trx     repository.ITTransactionRepo
// }

// type IEarningStatUpdUsecase interface {
// 	UpdateEarningStatus([]byte) []byte
// }

// func NewEarningStatUpdUsecase(earning repository.ITEarningRepo, trx repository.ITTransactionRepo) IEarningStatUpdUsecase {
// 	return &earningStatUpdUsecase{earning: earning, trx: trx}
// }

// func (u *earningStatUpdUsecase) UpdateEarningStatus(msg []byte) []byte {
// 	req := new(model.EarnPubRes)
// 	// if err := kafka.ExtractMsg(msg, &req); err != nil {
// 	// 	return kafka.Response(new(model.BaseResp).Err(err))
// 	// }
// 	if err := json.Unmarshal(msg, &req); err != nil {
// 		return kafka.Response(new(model.BaseResp).Err(err))
// 	}

// 	// create t_earning entity
// 	earnStatus := entity.TEarning{}
// 	copier.Copy(&earnStatus, req)

// 	// get trx by invoice no (reference id)
// 	trx, err := u.trx.FindOne(entity.TTransaction{InvoiceNo: req.ReferenceId, InstitutionId: req.InstitutionId})
// 	if err != nil {
// 		msg := fmt.Sprintf("ERROR: Finding transaction to update with invoice number %s and institution_id %s error : %s",
// 			req.ReferenceId, req.InstitutionId, err.Error())
// 		log.Println(msg)
// 		return kafka.Response(new(model.BaseResp).Err(errors.New(msg)))
// 	}

// 	// prepare query
// 	qry := make(map[string](string))
// 	qry["t_transaction_id"] = trx.BaseEntity.Id.String()

// 	// update status
// 	err = u.earning.Update(&earnStatus, qry)
// 	if err != nil {
// 		msg := fmt.Sprintf("ERROR: Updating earning data for transaction with invoice number %s and institution_id %s error : %s",
// 			req.ReferenceId, req.InstitutionId, err.Error())
// 		log.Println(msg)
// 		return kafka.Response(new(model.BaseResp).Err(errors.New(msg)))
// 	}

// 	return kafka.Response(new(model.BaseResp).OK())
// }
