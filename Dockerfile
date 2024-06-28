FROM golang:1.22

RUN go install github.com/cosmtrek/air@v1.49.0

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

CMD ["air", "-c", ".air.toml"]