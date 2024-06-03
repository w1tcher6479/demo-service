package cache

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/w1tcher6479/demo-service/internal/models"
)

func LoadCache(db *sql.DB, cacheData map[string]models.Order) {
	rows, err := db.Query(`SELECT * FROM orders`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var orderUID string
		var jsonStr string
		var order = models.Order{}

		err = rows.Scan(&orderUID, &jsonStr)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal([]byte(jsonStr), &order)
		if err != nil {
			log.Fatal(err)
		}
		cacheData[orderUID] = order
	}
}
