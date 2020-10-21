FROM golang:1.14.10

MAINTAINER Rodrigo Soares<rodrigo.escola@gmail.com>

RUN go get -u github.com/gin-gonic/gin

COPY app /go/

EXPOSE 8080

RUN go build -o app

CMD ./app


