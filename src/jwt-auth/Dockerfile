FROM golang:1.17-alpine as builder
WORKDIR /app

COPY . .

RUN go build main.go

FROM alpine:3.14

RUN apk update \
    && apk upgrade \
    && apk add --no-cache tzdata curl

ENV TZ Europe/Istanbul

WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
EXPOSE 3000