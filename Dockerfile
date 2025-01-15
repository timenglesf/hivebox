FROM --platform=$BUILDPLATFORM golang:1.23.4-alpine3.21 AS builder

RUN apk add --no-cache git=2.43.0-r0 make=4.4.1-r2
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o hivebox ./cmd/api

FROM alpine:3.17.6

WORKDIR /app

COPY --from=builder /app/hivebox /hivebox

EXPOSE 8282

ENTRYPOINT ["./hivebox"]
