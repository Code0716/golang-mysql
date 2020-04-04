FROM golang:alpine

WORKDIR /web

COPY /mysql /mysql

RUN apk update && \
    apk add --no-cache git && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/gin-gonic/gin

CMD ["go", "run", "main.go"] 