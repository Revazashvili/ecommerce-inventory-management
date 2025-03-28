package consumers

import (
	"context"
	"encoding/json"
	"log"

	pd "github.com/Revazashvili/ecommerce-inventory-management/product/database"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

const (
	broker = "localhost:29092"
)

var topics = []string{"products.ProductAdded", "products.ProductNameUpdated"}

func ListenToProductEvents(ctx context.Context, q *pd.Queries) {
	go listenForEvent(ctx, q)
}

func listenForEvent(ctx context.Context, q *pd.Queries) {
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

			var p ProductEvent

			err = json.Unmarshal(message.Value, &p)

			if err != nil {
				log.Printf("Error unmarshaling message: %s\n", err)
				continue
			}

			if *message.TopicPartition.Topic == "products.ProductAdded" {
				err := q.Insert(ctx, pd.InsertParams(pd.InsertParams{
					ID:   p.ProductId,
					Name: p.ProductName,
				}))
				if err != nil {
					log.Println(err)
					continue
				}
			} else {
				err := q.Update(ctx, pd.UpdateParams{ID: p.ProductId, Name: p.ProductName})
				if err != nil {
					log.Println(err)
					continue
				}
			}

		}
	}
}

type ProductEvent struct {
	ProductId   uuid.UUID
	ProductName string
}
