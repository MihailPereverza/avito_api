FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .
ADD go.sum .
ADD .env .
COPY cmd cmd
COPY internal internal
COPY docs docs

RUN go build -o avito_api cmd/http/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/avito_api /build/avito_api
ADD .env .
RUN mkdir -p ./reports

CMD ["./avito_api"]