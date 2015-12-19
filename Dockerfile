FROM golang:alpine

COPY . /go/src/github.com/wouterd/go-hello-web

RUN go install github.com/wouterd/go-hello-web

ENTRYPOINT /go/bin/go-hello-web

EXPOSE 8080