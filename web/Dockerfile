FROM golang:1.12.4-alpine

WORKDIR /go/src/app
COPY . /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache git \
    && go get github.com/oxequa/realize \
    && go build -o /go/src/app/goauth

CMD ["/go/src/app/goauth"]
