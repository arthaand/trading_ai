package earning

// import (
// 	earning "trading-ai/usecase/earning"

// 	"github.com/Shopify/sarama"
// )

// type earningHandler struct {
// 	consumer sarama.Consumer
// 	producer sarama.AsyncProducer
// 	uc       earning.IEarningStatUpdUsecase
// }

// // func NewEarningHandler(consumer sarama.Consumer,
// // 	producer sarama.AsyncProducer,
// // 	uc earning.IEarningStatUpdUsecase) *earningHandler {
// // 	return &earningHandler{consumer: consumer,
// // 		producer: producer,
// // 		uc:       uc}
// // }

// // func (h *earningHandler) EarningStatusUpdate(reqTopic, resTopic string) {
// // 	consumer := kafka.Consume(reqTopic, h.consumer)

// // 	// running continous service
// // 	go func() {
// // 		for msg := range consumer {
// // 			h.producer.Input() <- &sarama.ProducerMessage{
// // 				Topic: resTopic,
// // 				Value: sarama.ByteEncoder(h.uc.UpdateEarningStatus(msg.Value)),
// // 			}
// // 		}
// // 	}()
// // }
