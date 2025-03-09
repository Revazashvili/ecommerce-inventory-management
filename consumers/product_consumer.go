package consumers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Revazashvili/ecommerce-inventory-management/product"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	broker = "localhost:29092"
)

var topics = []string{"products.ProductAddedEvent", "products.ProductNameUpdatedEvent"}

func ListenToProductEvents(ctx context.Context, storage product.Storage) {
	go listenForEvent(ctx, storage)
}

func listenForEvent(ctx context.Context, storage product.Storage) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "products",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Printf("Failed to create consumer: %s\n", err)
		return
	}
	defer c.Close()

	err = c.SubscribeTopics(topics, nil)

	if err != nil {
		log.Printf("Failed to subscribe to topics: %s\n", err)
		return
	}

	log.Printf("Subscribed to topics: %s\n", topics)

	for {
		select {
		case <-ctx.Done():
			log.Printf("done")
			return
		default:
			message, err := c.ReadMessage(-1)

			if err != nil {
				log.Printf("Error reading message: %s\n", err)
				continue
			}

			var p product.Product

			err = json.Unmarshal(message.Value, &p)

			if err != nil {
				log.Printf("Error unmarshaling message: %s\n", err)
				continue
			}

			if *message.TopicPartition.Topic == "products.ProductAddedEvent" {
				_, err := storage.Add(ctx, p)
				if err != nil {
					log.Println(err)
					continue
				}
			} else {
				_, err := storage.Update(ctx, p)
				if err != nil {
					log.Println(err)
					continue
				}
			}

		}
	}
}
