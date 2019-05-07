FROM golang

WORKDIR $GOPATH/src/github.com/iaronaraujo/RedCoins

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 4000

CMD ["RedCoins"]