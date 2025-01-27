b:
	./mvnw package && docker build -f src/main/docker/Dockerfile.jvm -t quarkus/quarkus-vs-go-jvm . && \
	./mvnw package -Dquarkus.package.jar.type=legacy-jar && docker build -f src/main/docker/Dockerfile.legacy-jar -t quarkus/quarkus-vs-go-legacy-jar . && \
	./mvnw package -Dnative -Dquarkus.native.container-build=true && docker build -f src/main/docker/Dockerfile.native -t quarkus/quarkus-vs-go-native . && \
	docker build -f src/main/docker/Dockerfile.native-micro -t quarkus/quarkus-vs-go .

jvm:
	docker-compose up -d jvm

legacy-jar:
	docker-compose up -d jvm-legacy

native:
	docker-compose up -d native

micro:
	docker-compose up -d native-micro


down:
	docker-compose down