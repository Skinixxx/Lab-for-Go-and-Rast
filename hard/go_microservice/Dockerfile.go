FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go-service/go.mod go-service/go.sum ./
COPY go-service/proto/ ./proto/
RUN go mod download

COPY go-service/*.go ./
RUN go generate ./proto/
RUN CGO_ENABLED=0 GOOS=linux go build -o service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/service .
EXPOSE 8080 50051
CMD ["./service"]
