package subscriber

import (
	"encoding/json"
	"task-level-0/internal/domain/model"
	"task-level-0/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	service *service.Service
	conn    stan.Conn
	sub     stan.Subscription
	subject string
}

func NewSubscriber(conn stan.Conn, subj string, serv *service.Service) *Subscriber {
	return &Subscriber{
		service: serv,
		subject: subj,
		conn:    conn,
	}
}

func (s *Subscriber) Run() {
	vld := validator.New()
	subscription, err := s.conn.Subscribe(s.subject, func(msg *stan.Msg) {

		var order model.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			logrus.WithError(err).Info("subscriber: error with unmarshal msg data")
			return
		}

		if err := vld.Struct(order); err != nil {
			logrus.WithError(err).Info("subscriber: eror with validate msg data")
			return
		}

		id, err := s.service.AddOrder(order.OrderUID, msg.Data)
		if err != nil {
			logrus.WithError(err).Info("error occured with adding order into database")
			return
		}

		logrus.Infof("order added into database with id = %s", id)

	}, stan.DurableName("subscriber"))
	// sub.Unsubscribe()
	// s.conn.Close()
	if err != nil {
		logrus.Fatalf("subscriber: error occurred while connecting to nats-streaming")
	}

	s.sub = subscription

}

func (s *Subscriber) Stop() {
	s.sub.Unsubscribe()
	s.sub.Close()
}

// func PrettyString(str string) (string, error) {
// 	var prettyJSON bytes.Buffer
// 	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
// 		return "", err
// 	}
// 	return prettyJSON.String(), nil
// }
