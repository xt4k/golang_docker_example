FROM golang:alpine

WORKDIR /

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go","run","main.go"]