FROM golang:1.14

WORKDIR /go/src/app

COPY ./main.go .
COPY ./sql.go .
COPY ./templates.go .
COPY ./.env .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]