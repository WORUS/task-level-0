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
	status  bool
}

func NewPublisher(conn stan.Conn, subj string, stat bool) *Publisher {
	return &Publisher{
		conn:    conn,
		subject: subj,
		status:  stat,
	}
}

func (p *Publisher) Run() {
	//count := 0
	for {
		if !p.status {
			return
		}

		order := new(model.Order)
		order = generateJSON(order)

		b, err := json.Marshal(order)
		if err != nil {
			logrus.Fatal(err)
		}

		p.conn.Publish(p.subject, b)

		time.Sleep(1 * time.Second)
	}

}

func (p *Publisher) Stop() {
	p.status = false
}
