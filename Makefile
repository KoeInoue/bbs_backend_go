up:
	docker-compose up -d
down:
	docker-compose down --remove-orphans
exec: 
	docker container exec -it go_web ash