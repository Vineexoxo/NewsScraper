package rabbitmq

import (
	"context"
	"fmt"
	"time"
	"github.com/streadway/amqp"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
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

func NewRabbitMQConn(cfg *RabbitMQConfig, ctx context.Context, log logger.ILogger) (*amqp.Connection, error){
	if(cfg == nil){
		return nil, fmt.Errorf("rabbitMQconfig is nil")
	}
	
	connAddr:= fmt.Sprintf("amqp://%s:%s@localhost:%d/", cfg.User, cfg.Password, cfg.Port)

	bkoff:= backoff.NewExponentialBackOff()
	bkoff.MaxElapsedTime= 10 * time.Second
	maxRetries:=5
	fmt.Println("Connection information: ", connAddr)
	var conn *amqp.Connection
	var err error
	for i:=0; i<maxRetries; i++{
		conn, err= amqp.Dial(connAddr)
		if err!=nil{
			log.Errorf("Failed to connect to RabbitMQ: Connection information: %s",  connAddr, ctx)
			return conn, nil
		}
		time.Sleep(bkoff.NextBackOff())
	
	
	}
	fmt.Print("Connected to rabbmitMq")
	go func(){
		select {
			// context done : i.e connection is closed, so a goroutine that is always running 
			// and checking for whether the context is done or not 
			case <-ctx.Done():
				err:=conn.Close()
				if err!=nil {
					log.Errorf("Failed to close connection to RabbitMQ: %v", err, ctx)
				}else{
					log.Infof("Closed connection to RabbitMQ", ctx)
				}
		}
	}()
	return conn, err
}
