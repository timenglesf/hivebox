FROM golang:1.23.4-alpine3.21 as BUILDER

WORKDIR /app

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hivebox ./cmd/api

FROM scratch

COPY --from=builder /app/hivebox /hivebox

EXPOSE 8282

ENTRYPOINT ["/hivebox"]

