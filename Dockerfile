FROM golang:1.18

WORKDIR /myapp

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /emailsending

EXPOSE 8080

CMD ["/emailsending"]
