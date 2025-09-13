.PHONY: swag
swag:
	swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: build
build:
	go build -o bin/api cmd/api/main.go

.PHONY: start
start:
	swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
	go run cmd/api/main.go

.PHONY: serve
serve:
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
	@echo "Starting application..."
	@go run cmd/api/main.go