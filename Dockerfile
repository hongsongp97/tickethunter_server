FROM golang:1.9

WORKDIR $GOPATH/src/github.com/hongsongp97/tickethunter_server
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80