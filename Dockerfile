FROM golang:latest
WORKDIR $GOPATH/src/github.com/NotVasil/Password-Manager

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["pwdm"]