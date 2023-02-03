postgres:
	docker run --name goBinance -p 5444:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgresDB -d postgres

createdb:
# run on window
	docker exec -it goBinance bash -c "createdb --username=postgres --owner=postgres go_binance"

# run on macOS
# docker exec -it goBinance /bin/sh createdb --username=postgres --owner=postgres go_binance

dropdb:
# run on window
	docker exec -it goBinance bash -c "dropdb go_binance -U postgres"

# run on macOS
# docker exec -it goBinance /bin/sh dropdb go_binance -U postgres


migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5444/go_binance?sslmode=disable" -verbose up


migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5444/go_binance?sslmode=disable" -verbose down


sqlc: 
#  only run with cmd on window
	docker run --rm -v /C/Users/Acer/Desktop/go-labs/src/github.com/user/simplebank:/src -w /src kjconroy/sqlc generate

#  only run on ubuntu and macOS
# docker run --rm -v "C:\Users\Acer\Desktop\go-labs\src\github.com\user\simplebank:/src" -w /src kjconroy/sqlc generate
 
commit:
	git add .
	git commit -m"$t"
	git push origin main



server:
	go run main.go

devStart:
	CompileDaemon -command="./go-go_binance"

.PHONY: postgres createdb dropdb migratedown migrateup sqlc server commit devStart