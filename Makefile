compose: 
	docker compose up --build

run:
	go run ./cmd/service/

send:
	go run ./cmd/publisher/

clean:
	docker-compose down
