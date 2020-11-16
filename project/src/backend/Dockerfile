FROM golang:1.14.10

RUN go get -u github.com/lib/pq

EXPOSE 8081

COPY database.go /go/database.go

RUN go build -o app

CMD ./app