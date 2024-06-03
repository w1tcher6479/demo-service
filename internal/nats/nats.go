package nats

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"
	"github.com/w1tcher6479/demo-service/internal/database"
	"github.com/w1tcher6479/demo-service/internal/models"
)

func Connect(clusterID, clientID string) stan.Conn {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Не удалось подключиться к NATS Streaming: %v", err)
	}

	return sc
}

func Publish(sc stan.Conn, channel string, msg []byte) {
	if err := sc.Publish(channel, msg); err != nil {
		log.Fatalf("Не удалось опубликовать сообщение в канал: %v", err)
	}
}

func Subscribe(sc stan.Conn, db *sql.DB, channel string, cacheData map[string]models.Order) stan.Subscription {
	sub, err := sc.Subscribe(channel, func(msg *stan.Msg) {
		data := models.Order{}
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Fatal(err)
		}

		database.InsertOrder(db, data, msg.Data)

		log.Println("Received data through", msg.Subject)
		cacheData[data.OrderUID] = data
	})
	if err != nil {
		log.Fatalf("Не удалось подписаться на канал: %v", err)
	}
	return sub
}
