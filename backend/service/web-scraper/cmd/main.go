package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel"
)

// TestMessage is the message structure
type TestMessage struct {
	Text string `json:"text"`
}

// Handler for consuming TestMessage
func HandleTestMessage(queue string, delivery amqp.Delivery, _ *struct{}) error {
	fmt.Printf("📨 Received message: %s\n", string(delivery.Body))

	// Unmarshal to compare with sent message
	var receivedMsg TestMessage
	if err := json.Unmarshal(delivery.Body, &receivedMsg); err != nil {
		fmt.Printf("❌ Error unmarshalling message: %v\n", err)
	}

	// Ack the message
	if err := delivery.Ack(false); err != nil {
		fmt.Printf("❌ Error acknowledging message: %v\n", err)
	}

	fmt.Println("✅ Message acknowledged and removed from queue")
	return nil
}

func main() {
	logCfg := &logger.LoggerConfig{LogLevel: "debug"}
	log := logger.InitLogger(logCfg)

	cfg := &rabbitmq.RabbitMQConfig{
		Host: "localhost",
		Port: 5672,
		User: "guest",
		Password: "guest",
		Kind: "direct",
	}

	ctx := context.Background()

	// Connect to RabbitMQ
	conn, err := rabbitmq.NewRabbitMQConn(cfg, ctx, log)
	if err != nil {
		log.Errorf("❌ Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	tracer := otel.Tracer("test-publisher")

	// Start consumer first
	go func() {
		consumer := rabbitmq.NewConsumer[*struct{}](ctx, cfg, conn, log, tracer, HandleTestMessage)
		if err := consumer.ConsumeMessage(TestMessage{}, nil); err != nil {
			log.Errorf("❌ Error consuming message: %v", err)
		}
	}()

	// Give the consumer time to bind to the queue
	time.Sleep(1 * time.Second)

	// Create publisher
	publisher := rabbitmq.NewPublisher(ctx, cfg, conn, log, tracer)

	// Publish a test message
	msg := &TestMessage{Text: fmt.Sprintf("Hello RabbitMQ! Sent at %v", time.Now())}
	msgBytes, _ := json.Marshal(msg)

	if err := publisher.PublishMessage(msg); err != nil {
		log.Errorf("❌ Error publishing test message: %v", err)
		return
	}
	log.Infof("✅ Test message published successfully: %s", string(msgBytes))

	// Give time for consumer to process the message
	time.Sleep(2 * time.Second)

	// Verify queue is empty
	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("❌ Failed to open channel: %v", err)
		return
	}
	defer ch.Close()

	qName := "test_message_queue"
	queueInfo, err := ch.QueueInspect(qName)
	if err != nil {
		log.Errorf("❌ Failed to inspect queue: %v", err)
		return
	}

	if queueInfo.Messages == 0 {
		log.Info("✅ Queue is empty after consuming the message")
	} else {
		log.Infof("⚠️ Queue still has %d messages", queueInfo.Messages)
	}

	select {} // Keep app running to listen for any more messages
}
