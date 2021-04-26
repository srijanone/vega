# Builder stage
FROM golang:1.16.3-alpine AS builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go/bin/vega .


# Deploy stage
FROM alpine:latest

RUN apk update
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache git

COPY --from=builder /go/bin/vega .

ENTRYPOINT ["vega"]