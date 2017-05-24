FROM golang:1.8
RUN mkdir -p /go/src/lista/api/
COPY src /go/src/lista/api/
WORKDIR /go/src/lista/api/

RUN go build -o app main.go
CMD ["./app"]
