FROM golang:alpine

WORKDIR /app/cmd

COPY /mysql /mysql

RUN apk update && \
    apk add --no-cache git && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/gin-gonic/gin && \
    go get github.com/jinzhu/gorm && \
    go get github.com/gin-contrib/cors 
# && \
# go get github.com/rwcarlsen/goexif/exif
