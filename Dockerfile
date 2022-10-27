# Build Stage
FROM golang:1.19.2-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L "https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz" | tar xvz


# Run Stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY --from=builder /app/app.env .
COPY --from=builder /app/db/migration ./migration
COPY start.sh .
COPY wait-for.sh .

EXPOSE 8081
CMD ["/app/main"]
ENTRYPOINT [ "app/start.sh"]