FROM golang:1.14.10

RUN go get -u github.com/lib/pq
RUN go get gorm.io/gorm
RUN go get github.com/go-gorm/postgres
RUN go get -u github.com/gin-gonic/gin

EXPOSE 8081

COPY backend.go /go/backend.go

RUN go build -o app

CMD ./app