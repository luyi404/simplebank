# Build Stage
FROM golang:1.19.2-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run Stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8081
CMD ["/app/main"]