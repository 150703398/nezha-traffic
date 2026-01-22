# --------- Build Stage ---------
FROM golang:1.22-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git build-base sqlite-dev bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o nezha-dashboard

# --------- Runtime Stage ---------
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata bash sqlite

WORKDIR /dashboard
COPY --from=builder /app/nezha-dashboard /dashboard/nezha-dashboard
COPY docker/entrypoint.sh /entrypoint.sh
COPY resource /dashboard/resource

RUN chmod +x /entrypoint.sh
EXPOSE 8008
ENTRYPOINT ["/entrypoint.sh"]
