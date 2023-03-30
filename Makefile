.PHONY: build
build:
	go build -o bin/goshorten main.go

.PHONY: run
run: build
	@./bin/goshorten

.PHONY: dev
dev:
	@go run main.go

.PHONY: test
test:
	@go test -v ./...
