postgres:
	docker run --name beacon-indexer-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it beacon-indexer-postgres createdb --username=root --owner=root beacon-indexer

dropdb:
	docker exec -it beacon-indexer-postgres dropdb beacon-indexer

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/beacon-indexer?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/beacon-indexer?sslmode=disable" -verbose down