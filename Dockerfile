FROM golang:1.19

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 80

CMD ["air", "-c", ".air.toml"]
