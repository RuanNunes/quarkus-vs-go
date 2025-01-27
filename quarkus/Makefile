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

load-native-micro:
	wrk -t10 -c10 -d1m http://0.0.0.0:8080/hello
load-native:
	wrk -t10 -c10 -d1m http://0.0.0.0:8081/hello
load-jvm-legacy:
	wrk -t10 -c10 -d1m http://0.0.0.0:8082/hello
load-jvm:
	wrk -t10 -c10 -d1m http://0.0.0.0:8083/hello


down:
	docker-compose down