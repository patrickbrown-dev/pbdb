FROM golang:1.16-alpine

WORKDIR /pbdb
COPY . .

RUN apk add git
RUN go build -v ./...
RUN go install -v ./...
RUN mkdir -p /etc/pbdb/

EXPOSE 1728:1728

CMD ["pbdb", "run"]