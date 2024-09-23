FROM golang:1.23

RUN go install github.com/air-verse/air@latest

WORKDIR /app

CMD ["air", "-c", ".air.toml"]