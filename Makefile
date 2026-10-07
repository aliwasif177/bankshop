postgres:
	docker run --name postgres --network bank-network \
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

migrateup1:
	migrate -path db/migration \
	-database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
	up 1

migratedown:
	migrate -path db/migration \
	-database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
	down

migratedown1:
	migrate -path db/migration \
	-database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
	down 1

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/aliwasif177/bankshop/db/sqlc Store
	
test:
	go test -v -cover ./...

server:
	go run main.go

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock