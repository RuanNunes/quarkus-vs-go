go-up-b:
	cd go && docker-compose up -d --build
go-up:
	cd go && docker-compose up -d

quarkus-up-b:
	cd quarkus &&  make b && docker-compose up -d
quarkus-up:
	cd quarkus && docker-compose up -d

b-up:
	make go-up-b  && make quarkus-up-b
up:
	make go-up  && make quarkus-up
down:
	cd go && docker-compose down
	cd quarkus && docker-compose down