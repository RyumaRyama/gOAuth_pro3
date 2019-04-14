FROM golang:1.12.4

RUN go get github.com/labstack/echo/...
RUN go get -u gopkg.in/godo.v2/cmd/godo

WORKDIR /app
COPY . /app

CMD ["go", "run", "main.go"]
