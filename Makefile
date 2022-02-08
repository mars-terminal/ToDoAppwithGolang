.phony: migrate_down_postgres migrate_up_postgres

export POSTGRESQL_URL= "postgres://rustam:1234@192.168.1.138:5432/my_world?sslmode=disable"

migrate_up_postgres:
	migrate -database ${POSTGRESQL_URL} -path schema up
migrate_down_postgres:
	migrate -database ${POSTGRESQL_URL} -path schema down

run:
	go run cmd/app/main.go
