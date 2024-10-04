up:
	@docker-compose up 
up-no-cache:
	@docker-compose up --build  
down:
	@docker-compose down
front-bash:
	@docker exec -it vue-go_web_1 bash
back-bash:
	@docker exec -it vue-go_go-app_1 sh
local-prod:

	@docker build -t vue-go . && docker run -p 8080:8080 -e APP_NAME=SIR -e SERVER_HOST=localhost -e SERVER_PORT=8080 -e DB_USERNAME=freedb_gobabak -e DB_PASSWORD=AyaRn8%ymVPH!uA -e DB_HOST=sql.freedb.tech -e DB_PORT=3306 -e DB_NAME=freedb_gobabak vue-go

test-prod:
	@docker run --publish 8080:8080 bablido/my-docker-image:latest -e APP_NAME=SIR -e SERVER_HOST=localhost -e SERVER_PORT=8080 -e DB_USERNAME=freedb_gobabak -e DB_PASSWORD=AyaRn8%ymVPH!uA -e DB_HOST=sql.freedb.tech -e DB_PORT=3306 -e DB_NAME=freedb_gobabak vue-go

.PHONY: up down front-bash up-no-cache local-prod test-prod



