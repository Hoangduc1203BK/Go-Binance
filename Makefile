postgres:u
	docker run --name goBinance up 5444:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgresDB -d postgres

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


commit:
	git add .
	git commit -m"$t"
	git push origin main

pull:
	git pull origin main

server:
	go run main.go

devStart:
	CompileDaemon -command="./Go-Binance"

.PHONY: postgres createdb dropdb migrateup sqlc server commit devStart