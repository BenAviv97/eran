# Build stage
FROM golang:1.22 as builder
WORKDIR /src
COPY ../../../.. .
RUN CGO_ENABLED=0 go build -o /app/api-gateway ./livestreampro/backend/cmd/api-gateway

# Final stage
FROM alpine:3.18
COPY --from=builder /app/api-gateway /api-gateway
ENTRYPOINT ["/api-gateway"]
