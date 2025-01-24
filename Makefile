up:
	swag init
	@docker compose up -d && sleep 3 && docker ps -a

down:
	@docker compose down