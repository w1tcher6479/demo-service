package main

import (
	"github.com/w1tcher6479/demo-service/config"
	"github.com/w1tcher6479/demo-service/internal/files"
	"github.com/w1tcher6479/demo-service/internal/nats"
)

func main() {
	jsonData := files.ReadFile("assets/model.json")

	sc := nats.Connect(config.ClusterID, config.PublisherID)
	defer sc.Close()

	nats.Publish(sc, config.ChannelName, jsonData)
}
