FROM golang:1.18

RUN go install gopkg.in/yaml.v2@latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main

EXPOSE 8080

CMD ["/app/main"]