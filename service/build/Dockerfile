FROM golang:1.14-alpine3.11 AS build-env
WORKDIR /go/src/tratnik.net/service
COPY . /go/src/tratnik.net/service
RUN go build -o /server

FROM alpine:3.11
WORKDIR /opt/service
COPY ./migrations ./migrations
COPY ./config/dev.yaml ./config.yaml
COPY --from=build-env /server .
ENTRYPOINT ["./server"]
