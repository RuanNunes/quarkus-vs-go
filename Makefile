go-up:
	cd go && docker-compose up -d --build

quarkus-up:
	cd quarkus &&  make b && docker-compose up -d


up:
	make go-up && make quarkus-up

down:
	cd go && docker-compose down
	cd quarkus && docker-compose down