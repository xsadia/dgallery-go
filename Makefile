-include .env

DB_URL=$(POSTGRES_URL)

test:
	go test -v ./...

migration-create:
	migrate create -ext sql -dir ./migrations -seq $(name)

migrate-up:
	migrate -path migrations/ -database $(DB_URL) -verbose up

migrate-down:
	migrate -path migrations/ -database $(DB_URL) -verbose down $(times)

compose-up:
	sudo docker-compose up $(flag)

compose-down:
	sudo docker-compose down
