# Build stage
FROM --platform=$BUILDPLATFORM golang:1.24.0-alpine3.21 AS builder

RUN apk update && \
  apk add --no-cache git=2.47.2-r0 make=4.4.1-r2
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application with platform-specific settings
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT=""
RUN echo "Building for $TARGETOS/$TARGETARCH" && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT#v} \
  go build -o hivebox ./cmd/api


# Final stage
FROM --platform=$TARGETPLATFORM alpine:3.21@sha256:56fa17d2a7e7f168a043a2712e63aed1f8543aeafdcee47c58dcffe38ed51099

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/hivebox .

EXPOSE 8282

ENTRYPOINT ["./hivebox"]
