swagger-models:
	swagger generate model \
  	--spec=docs/swagger.yml \
  	--model-package=internal/types

migrate-up:
	sql-migrate up

migrate-down:
	sql-migrate down

migrate-reset:
	sql-migrate reset

fetch-models:
	sqlboiler psql

setup-db:
	sql-migrate up && sqlboiler psql

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

serve:
	make migrate-up && make docs && make run