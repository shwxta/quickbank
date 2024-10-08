postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root quickbank

dropdb:
	docker exec -it postgres dropdb quickbank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/quickbank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/quickbank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server