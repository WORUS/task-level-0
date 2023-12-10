package publisher

import (
	"encoding/json"
	"task-level-0/internal/domain/model"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	conn    stan.Conn
	subject string
}

func NewPublisher(conn stan.Conn, subj string) *Publisher {
	return &Publisher{
		conn:    conn,
		subject: subj,
	}
}

func (p *Publisher) Run() {

	for {
		order := new(model.Order)
		order = generateJSON(order)

		b, err := json.Marshal(order)
		if err != nil {
			logrus.Fatal(err)
		}
		p.conn.Publish(p.subject, b)
		time.Sleep(2 * time.Minute)
	}
}
