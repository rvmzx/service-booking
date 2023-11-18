FROM golang:1.20

WORKDIR /go/src/app

COPY . .

RUN go get -u github.com/gorilla/mux
RUN go get -u go.mongodb.org/mongo-driver/mongo

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]