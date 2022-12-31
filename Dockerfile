FROM golang:1.19-alpine3.15 AS build

WORKDIR  /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/benchmark-migrations ./infra/db/migrations/...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/benchmark-query ./cmd/...

FROM alpine:3.13.6

WORKDIR /timescaledb

COPY --from=build /build/benchmark-migrations /usr/bin/
COPY --from=build /build/benchmark-query /usr/bin/
COPY --from=build /build/migrations /timescaledb/migrations
COPY --from=build /build/tests/data /timescaledb