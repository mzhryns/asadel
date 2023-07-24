.PHONY: init
init:
	go mod download

.PHONY: launch
launch:
	go run main.go