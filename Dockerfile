FROM golang:1.14

WORKDIR /app

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    | sh -s -- -b $(go env GOPATH)/bin

CMD $(go env GOPATH)/bin/air
