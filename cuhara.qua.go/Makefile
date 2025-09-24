.PHONY: docs
docs:
	@echo "Generating Swagger documentation..."
	swagger generate model \
  	--spec=docs/swagger.yml \
  	--model-package=internal/types
	@echo "Swagger documentation generated successfully"

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: build
build:
	go build -o bin/api cmd/api/main.go

.PHONY: serve
serve:
	make docs && make run