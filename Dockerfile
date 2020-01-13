FROM golang:1.13.6

RUN mkdir /go/src/grpc_test

ADD . /go/src/grpc_test

WORKDIR /go/src/grpc_test

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN go build .

CMD ["./grpc_test"]
