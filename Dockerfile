# Builder stage
FROM golang:1.14.1-alpine AS builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go/bin/vega .


# Deploy stage
FROM alpine:3.15

RUN apk update
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache git

COPY --from=builder /go/bin/vega .

ENTRYPOINT ["vega"]