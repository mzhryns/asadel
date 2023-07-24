.PHONY: init
init:
	go mod download

.PHONY: launch
launch:
	go run main.go

.PHONY: docker-build
docker-build:
	docker build -t asadel .

.PHONY: docker-run
docker-run:
	docker-compose up -d
