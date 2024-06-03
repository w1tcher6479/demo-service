package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/w1tcher6479/demo-service/config"
	"github.com/w1tcher6479/demo-service/internal/cache"
	"github.com/w1tcher6479/demo-service/internal/database"
	"github.com/w1tcher6479/demo-service/internal/handlers"
	"github.com/w1tcher6479/demo-service/internal/models"
	"github.com/w1tcher6479/demo-service/internal/nats"
)

var (
	cacheData = make(map[string]models.Order, 10)
)

func main() {
	db := database.InitDB()
	defer db.Close()

	cache.LoadCache(db, cacheData)

	sc := nats.Connect(config.ClusterID, config.ServiceID)
	defer sc.Close()

	sub := nats.Subscribe(sc, db, config.ChannelName, cacheData)
	defer sub.Unsubscribe()

	time.Sleep(1 * time.Second)

	server := http.NewServeMux()
	server.HandleFunc("GET /order/{id}", handlers.GetOrder(cacheData))
	log.Fatal(http.ListenAndServe(":8080", server))
}
