FROM golang:1.9

WORKDIR /go
COPY . .

RUN go get -d -v ./...
RUN go build frontend.go

EXPOSE 3000
ENTRYPOINT ["/go/frontend"]
