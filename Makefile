.PHONY: init
init:
	go mod download

.PHONY: launch
launch:
	go run main.go

.PHONY: docker-build
launch:
	docker build -t asadel .

.PHONY: docker-run
launch:
	docker-compose up -d
