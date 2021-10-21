FROM golang:1.17-stretch

ENV GO111MODULE=on

# Setup document root
RUN mkdir -p /app/excell/code
RUN mkdir -p /app/excell/logs
RUN mkdir -p /app/excell/k6
RUN mkdir -p /app/excell/tmp

WORKDIR /app/excell/code

COPY go.mod /app/excell/code
COPY go.sum /app/excell/code
COPY reflex.conf /app/excell/code

RUN go mod download

RUN ["go", "install", "github.com/cespare/reflex@latest"]

COPY src src
COPY docker/env/excell-socket-local.env .env
COPY zgEXCELL-Socket.iml zgEXCELL-Socket.iml

EXPOSE 8080

ENTRYPOINT ["reflex", "-c", "./reflex.conf"]
