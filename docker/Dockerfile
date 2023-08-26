FROM golang:1.21.0-bullseye

WORKDIR /app

COPY . .

RUN go build -o srv ./server

EXPOSE 8080

CMD ["./srv"]
