DB_URL=postgresql://idagoras:314159@localhost:3306/adventure?sslmode=disable
postgres:
	docker run --name postgres16 -e POSTGRES_USER=idagoras -e POSTGRES_PASSWORD=314159 -p 3306:5432 -d postgres
createdb:
	docker exec -it postgres16 createdb --username=idagoras --owner=idagoras adventure
dropdb:
	docker exec -it postgres16 dropdb -U idagoras adventure
db_docs:
	dbdocs build doc/db.dbml
db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml
migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up
migrateup1:
	migrate -path db/migration -database "${DB_URL}" -verbose up 1
migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down
migratedown1:
	migrate -path db/migration -database "${DB_URL}" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./src/...
server:
	go run src/main/main.go
mock:
	mockgen -package mockdb -destination src/database/mock/store.go oceanus/src/database Store
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        proto/*.proto
evans:
	evans --host localhost --port 9090 -r repl

.PHONY:postgres createdb dropdb migratedown migrateup test server mock migratedown1 migrateup1 db_docs db_schema proto