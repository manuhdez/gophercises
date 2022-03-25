deps:
	go mod download
	go mod tidy
	go mod verify

.PHONY: start
start:
	@docker compose up -d

stop:
	@docker-compose stop

down:
	@docker-compose down -v --remove-orphans

status:
	@docker ps
