FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY cmd/main /app/
COPY pkg /app/pkg/
COPY go.mod /app/
COPY go.sum /app/

RUN apk --no-cache add bash git make gcc gettext musl-dev

RUN go mod download

RUN go build -o main main.go

FROM alpine:latest

COPY --from=builder /app/main /app/main

CMD ["/app/main"]

HEALTHCHECK --interval=5s --timeout=5s --start-period=5s --retries=3 CMD ["/app/healthcheck.sh"]

ENV PORT 8080
ENV HOST 0.0.0.0
ENV DB_HOST 0.0.0.0
ENV DB_PORT 5432
ENV DB_USER postgres
ENV DB_PASS postgres
ENV DB_NAME postgres
ENV DB_SCHEMA public
ENV DB_SSL_MODE disable