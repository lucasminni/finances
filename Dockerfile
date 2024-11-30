FROM golang:alpine

WORKDIR /app

RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./migrations /migrations

ENV DB_URL="postgres://user:admin@admin:5432/financas-db?sslmode=disable"

CMD ["migrate", "-path", "/migrations", "-database", "DB_URL", "up"]
