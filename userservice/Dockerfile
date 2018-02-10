FROM golang:1.9

WORKDIR /go
COPY . .

RUN go get -d -v ./...
RUN go build user-service.go

ENTRYPOINT ["/go/user-service"]
