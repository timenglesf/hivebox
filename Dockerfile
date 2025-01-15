FROM golang:1.23.4-alpine3.21 as builder 

RUN apk add --no-cache git make

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o hivebox ./cmd/api

FROM --platform=${TARGETOS}/${TARGETARCH} alpine:latest

WORKDIR /app

COPY --from=builder /app/hivebox /hivebox

EXPOSE 8282

ENTRYPOINT ["./hivebox"]

