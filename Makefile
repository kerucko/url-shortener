all: build docker-build docker-run

build: ./cmd/main.go
	go build -o ./build/app ./cmd/main.go

docker-build: docker-compose.yml
	docker compose build

docker-run: docker-compose.yml
	docker compose up

clean:
	rm ./build/app