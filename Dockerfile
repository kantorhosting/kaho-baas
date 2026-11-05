FROM golang:1.27-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Development: hot reload with air, source code is mounted from the host
FROM base AS dev

RUN go install github.com/air-verse/air@v1.61.7

COPY . .

CMD ["air", "-c", ".air.toml"]

FROM base AS build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/main ./cmd/api

# Production: static binary on a minimal image, running as non-root user
FROM alpine:3.20 AS prod

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app \
    && adduser -S -G app app

WORKDIR /app

COPY --from=build /out/main /app/main

USER app

EXPOSE 8080

CMD ["./main"]
