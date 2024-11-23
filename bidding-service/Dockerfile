FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY bidding-service/go.* ./
RUN go mod download
COPY bidding-service/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o bidding-service ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bidding-service .
EXPOSE 8086
CMD ["./bidding-service"]

