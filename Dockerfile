FROM golang:1.21 as build-env

WORKDIR /usr/local/go/src/international-trading-test


COPY . .

RUN go get ./...
RUN go build -o /international-trading-test international-trading-test/cmd/api

FROM ubuntu

COPY --from=build-env /international-trading-test /

CMD ["/international-trading-test"]

EXPOSE 8081