docker-up:
	@docker compose up -d && sleep 3 && docker ps -a

docker-down:
	@docker compose down