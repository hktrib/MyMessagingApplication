postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.4

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root my_messaging_app

dropdb:
	docker exec -it postgres15 dropdb my_messaging_app

dbschema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/my_messaging_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/my_messaging_app?sslmode=disable" -verbose down

sqlc:
	sqlc generate 

sqlc_delete:
	rm -rf ./db/sqlc

.PHONY: postgres createdb dropdb migrateup migratedown sqlc sqlc_delete