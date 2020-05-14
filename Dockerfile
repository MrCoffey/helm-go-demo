FROM golang:1.13.0-alpine

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN apk add git
RUN go get github.com/gorilla/mux
RUN go build -o main .

CMD ["/app/main"]
