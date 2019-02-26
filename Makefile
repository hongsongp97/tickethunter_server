start:
	docker-compose up -d

run:
	docker exec -it running_server tickethunter_server

bash: 
	docker exec -it running_server bash

stop:
	docker-compose down

restart:
	make stop
	make start

rebuild:
	docker-compose up -d --build