FROM golang:1.14

WORKDIR /go/src/app

COPY . .

RUN go build -o ./redis-go-example .

EXPOSE 8080

CMD [ "./redis-go-example" ]
