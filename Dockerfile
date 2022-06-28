FROM golang:1.18

RUN go install gopkg.in/yaml.v2

RUN mkdir /work

ADD . /work

WORKDIR /work

RUN go build -o app ./gui/app.go

EXPOSE 8080

CMD ["/work/gui/app"]