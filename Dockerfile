FROM golang:latest

WORKDIR $GOPATH/src/awesomePan
COPY . $GOPATH/src/awesomePan
RUN go get -u github.com/go-sql-driver/mysql
RUN go build main.go

EXPOSE 8088
ENTRYPOINT ["./awesomepan"]