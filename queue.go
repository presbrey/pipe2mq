package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type request struct {
	headers amqp.Table
	body    string
}
type requestCh chan *request

var (
	backlog  = flag.Uint("backlog", 8192, "incoming channel capacity")
	backoff  = flag.Duration("backoff", time.Second, "pause between errors")
	exchange = flag.String("exchange", "test", "AMQP exchange name")
	key      = flag.String("key", "test", "AMQP routing key")
	tag      = flag.String("tag", "pipe2mq", "AMQP consumer tag")
	uri      = flag.String("uri", "amqp://localhost:5672/", "AMQP URI")

	incoming = make(requestCh, *backlog)
)

func newRequest(body string) *request {
	return &request{
		headers: amqp.Table{
			"Time": fmt.Sprintf("%d", time.Now().Unix()),
		},
		body: body,
	}
}

func dial() (ch *amqp.Channel, err error) {
	conn, err := amqp.Dial(*uri)
	if err != nil {
		return
	}

	channel, err := conn.Channel()
	if err != nil {
		return
	}

	ch = channel
	return
}

func (rc requestCh) publish() error {
	var (
		err error
		out *amqp.Channel
	)
	for r := range rc {
		if r == nil {
			break
		}
		for out == nil {
			out, err = dial()
			if err != nil {
				out = nil
				log.Println(err)
				time.Sleep(*backoff)
			}
		}

		err = out.Publish(*exchange, *key, false, false,
			amqp.Publishing{
				Headers:      amqp.Table(r.headers),
				DeliveryMode: amqp.Persistent,
				Body:         []byte(r.body),
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	go func() {
		for {
			err := incoming.publish()
			if err != nil {
				log.Println(err)
			}
			time.Sleep(*backoff)
		}
	}()
}
