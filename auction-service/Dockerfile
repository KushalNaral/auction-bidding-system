FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY auction-service/go.* ./
RUN go mod download
COPY auction-service/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o auction-service ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/auction-service .
EXPOSE 8085
CMD ["./auction-service"]
