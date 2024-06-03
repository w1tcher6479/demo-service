package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/w1tcher6479/demo-service/config"
	"github.com/w1tcher6479/demo-service/internal/models"
)

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InsertOrder(db *sql.DB, data models.Order, jsonData []byte) {
	req := `
		INSERT INTO orders(order_uid, data)
		VALUES ($1, $2)
		RETURNING order_uid`

	var orderUID string

	err := db.QueryRow(req, data.OrderUID, jsonData).Scan(&orderUID)
	if err != nil {
		log.Println(err)
	}
}
