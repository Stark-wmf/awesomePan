FROM golang:latest

WORKDIR $GOPATH/src/awesomePan
COPY . $GOPATH/src/awesomePan
RUN go install github.com/go-sql-driver/mysql
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./awesomepan"]