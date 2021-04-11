FROM golang:1.14-alpine3.11 AS build-env
WORKDIR /go/src/tratnik.net/client
COPY . /go/src/tratnik.net/client
RUN go build -o /client

FROM alpine:3.11
WORKDIR /opt/client
COPY ./config/dev.yaml ./config.yaml
COPY --from=build-env /client .
CMD ["sleep", "infinity"]
