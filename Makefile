up:
	@docker-compose up 
up-no-cache:
	@docker-compose up --build  
down:
	@docker-compose down
front-bash:
	@docker exec -it vue-go_web_1 bash

.PHONY: up down front-bash up-no-cache