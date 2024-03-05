postgres:
	docker run --name postgres16 -e POSTGRES_USER=idagoras -e POSTGRES_PASSWORD=314159 -p 3306:5432 -d postgres
createdb:
	docker exec -it postgres16 createdb --username=idagoras --owner=idagoras adventure
dropdb:
	docker exec -it postgres16 dropdb -U idagoras adventure
migrateup:
	migrate -path db/migration -database "postgresql://idagoras:314159@localhost:3306/adventure?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://idagoras:314159@localhost:3306/adventure?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://idagoras:314159@localhost:3306/adventure?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://idagoras:314159@localhost:3306/adventure?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./src/...
server:
	go run src/main/main.go
mock:
	mockgen -package mockdb -destination src/database/mock/store.go bluesell/src/database Store
.PHONY:postgres createdb dropdb migratedown migrateup test server mock migratedown1 migrateup1