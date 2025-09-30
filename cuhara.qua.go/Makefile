swagger-models:
	oapi-codegen -config api/config/oapi-codegen.types.yml \
  	-o internal/types/openapi_types.gen.go \
  	api/swagger.yml

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