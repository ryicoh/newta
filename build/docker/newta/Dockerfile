FROM golang:1.13-alpine as builder

WORKDIR /build

COPY . .

RUN go build -o /newta ./cmd/newta/main.go

FROM alpine:3.11

COPY --from=builder /newta /usr/local/bin/newta

CMD newta
