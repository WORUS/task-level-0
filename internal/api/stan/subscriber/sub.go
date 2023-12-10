package subscriber

import (
	"bytes"
	"encoding/json"
	"task-level-0/internal/domain/model"
	"task-level-0/internal/service"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type OrderWriter interface {
	AddOrder()
}

type Subscriber struct {
	service *service.Service
	conn    stan.Conn
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
	_, err := s.conn.Subscribe(s.subject, func(msg *stan.Msg) {

		// res, err := PrettyString(string(msg.Data))
		// if err != nil {
		// 	logrus.Fatal("Error occurred while parse nats msg")
		// }
		// fmt.Println(res)
		//TODO: validate

		var order model.Order

		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			logrus.WithError(err).Info("Error with Unmarshal msg data")
		} else {
			id, err := s.service.AddOrder(order.OrderUID, msg.Data)
			if err != nil {
				logrus.WithError(err).Info("Error occured with adding order into database")
			} else {
				logrus.Infof("order added into database with id = %s", id)
			}
		}

	})
	// sub.Unsubscribe()
	// s.conn.Close()
	if err != nil {
		logrus.Fatalf("Error occurred while connecting to nats-streaming")
	}
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
