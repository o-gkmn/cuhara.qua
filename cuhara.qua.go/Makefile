docs:
	@echo "Generating Swagger documentation..."
	swagger generate model \
  	--spec=docs/swagger.yml \
  	--model-package=internal/types
	@echo "Swagger documentation generated successfully"

migrate-up:
	export $(cat .env | xargs) && sql-migrate up

migrate-down:
	export $(cat .env | xargs) && sql-migrate down

migrate-reset:
	export $(cat .env | xargs) && sql-migrate reset

fetch-models:
	export $(cat .env | xargs) && sqlboiler psql

setup-db:
	export $(cat .env | xargs) && sql-migrate up && sqlboiler psql

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

serve:
	make migrate-up && make docs && make run