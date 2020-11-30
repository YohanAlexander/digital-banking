ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

.PHONY: build
build:
	go build -o main cmd/banking/main.go

.PHONY: run
run:
	@docker-compose up -d

.PHONY: stop
stop:
	@docker-compose down

.PHONY: logs
logs:
	@docker-compose logs -f
