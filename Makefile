postgres:
	docker run --name postgres \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=secret \
	-e POSTGRES_DB=simple_bank \
	-p 5432:5432 \
	-d postgres:18-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration \
	-database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
	up

migratedown:
	migrate -path db/migration \
	-database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
	down

test:
	go test -v -cover ./...

server:
	go run main.go

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server