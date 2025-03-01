up:
	@swag init
	@docker compose up -d
	@docker ps -a

down:
	@docker compose down

test:
	@go clean -testcache
	@go test -coverprofile=coverage.out ./internal/.../
	@go tool cover -func=coverage.out
