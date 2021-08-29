FROM golang:1.16.7

WORKDIR /go/src/github.com/kenta4327/example-of-ec-site
COPY go.mod .
COPY go.sum .
RUN go install github.com/golang/mock/mockgen
