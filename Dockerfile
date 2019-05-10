FROM golang

WORKDIR $GOPATH/src/github.com/iaronaraujo/RedCoins

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 3000/tcp

CMD ["RedCoins"]