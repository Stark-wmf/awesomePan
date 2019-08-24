FROM golang:latest

WORKDIR $GOPATH/src/awesomePan
COPY . $GOPATH/src/awesomePan

RUN go build main.go

EXPOSE 8080
ENTRYPOINT ["./awesomepan"]