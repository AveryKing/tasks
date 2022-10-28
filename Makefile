postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=test -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root tasks

dropdb:
	docker exec -it postgres dropdb tasks

migrateup:
	migrate -path db/migration -database "postgresql://root:test@localhost:5432/tasks?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:test@localhost:5432/tasks?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
