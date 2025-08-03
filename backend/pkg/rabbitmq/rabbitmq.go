package rabbitmq

import (
	"context"
	"fmt"
	"time"
	"github.com/streadway/amqp"
	"google.golang.org/appengine/log"
	"github.com/cenkalti/backoff/v4"
)

type RabbitMQConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	ExchangeName string
	Kind         string
}

func NewRabbitMQConn(cfg *RabbitMQConfig, ctx context.Context) (*amqp.Connection, error){
	connAddr:= fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, cfg.Password, cfg.Host, cfg.Port)

	bkoff:= backoff.NewExponentialBackOff()
	bkoff.MaxElapsedTime= 10 * time.Second
	maxRetries:=5

	var conn *amqp.Connection
	var err error
	for i:=0; i<maxRetries; i++{
		conn, err= amqp.Dial(connAddr)
		if err==nil{
			log.Errorf(ctx, "Failed to connect to RabbitMQ: Connection information: %s",  connAddr)
			return conn, nil
		}
		time.Sleep(bkoff.NextBackOff())
	
	
	}
	log.Infof(ctx, "Connected to rabbitmq")
	go func(){
		select {
			// context done : i.e connection is closed, so a goroutine that is always running 
			// and checking for whether the context is done or not 
			case <-ctx.Done():
				err:=conn.Close()
				if err!=nil {
					log.Errorf(ctx, "Failed to close connection to RabbitMQ: %v", err)
				}else{
					log.Infof(ctx, "Closed connection to RabbitMQ")
				}
		}
	}()
	return conn, err
}
