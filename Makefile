ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o banking cmd/main.go

.PHONY: run
run:
	@docker-compose up -d --build

.PHONY: stop
stop:
	@docker-compose down

.PHONY: logs
logs:
	@docker-compose logs -f
