FROM golang:1.9

WORKDIR $GOPATH/src/github.com/hongsongp97/tickethunter_server
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80

# ENTRYPOINT [ "sh -c" ]

# CMD ["tickethunter_server"]
# CMD ["pwd"]
# CMD ["go", "run", "../temp/running.go"]
# CMD [ "cat" ]
# CMD [ "sleep infinity" ]

