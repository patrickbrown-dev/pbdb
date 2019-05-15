FROM golang:1.12-alpine

WORKDIR /go/src/pbdb
COPY . .

RUN apk add dep git
RUN dep ensure
RUN go install -v ./...
RUN mkdir -p /etc/pbdb/

EXPOSE 1728:1728

CMD [ "pbdb", "run"]