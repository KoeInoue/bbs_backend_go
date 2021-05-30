up:
	docker-compose up -d
down:
	docker-compose down --remove-orphans
web: 
	docker container exec -it go_web ash
db: 
	docker container exec -it mysql bash
logs:
	docker-compose logs -f web