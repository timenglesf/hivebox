FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY . /app

ENV HIVEBOX_VERSION="0.0.1" 

ENTRYPOINT [ "go", "run", "./cmd/api/main.go" ]
