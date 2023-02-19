FROM golang:1.20-bullseye as builder

WORKDIR /build
RUN apt-get install -y gcc

COPY go.* .
RUN go mod download

COPY . .
RUN go build -tags "json,vtable" -o movie ./cmd/server

FROM debian:bullseye as prod

WORKDIR /app
COPY --from=builder /build/movie movie
COPY database .

CMD "./movie"
