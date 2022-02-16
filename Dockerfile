FROM golang:1.17.0-alpine3.14

ENV PATH=/usr/local/git/bin:$PATH

RUN apk add --no-cache git

WORKDIR ${PATH}

COPY main.go ${ROOT}

COPY go.mod ${ROOT}

RUN go mod download

RUN go mod tidy