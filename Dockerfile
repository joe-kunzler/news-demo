# syntax=docker/dockerfile:1

FROM golang:1.15-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY . .
RUN go mod download

RUN ls -a

RUN pwd

RUN go clean -modcache

RUN go build

EXPOSE 8080

CMD [ "/app/news-demo-starter-files" ]