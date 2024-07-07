all: build

build: 
	@echo "Building..."
	@templ generate
	@go build -o ./bin/main cmd/api/main.go

run:
	@echo "Running..."
	@go run cmd/api/main.go

watch:
	@air -c .air.toml

dbmigrate: 
	@echo "Migrating..."
	TURSO_TOKEN=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MjAzNjgxMDYsImlkIjoiNGIxNDA3NDktN2RkYy00NmY1LWE4N2YtYmE4ZmJhOTkwZjVlIn0.h0gPgrKowrTp3zp9NUhdM7A0aZ6bkyWx2XevYhhFbx-8IIzN7gmMl8T9Z3giXgXEf5Og6fNDNMi_kXOXE81hAg atlas schema apply --env turso --to file://schema.sql --dev-url sqlite://dev.db
dbgenerate:
	@echo "Generating..."
	@sqlc generate
.PHONY: build run
