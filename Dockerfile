FROM golang:1.21

WORKDIR /app

RUN go mod init math

COPY . .

RUN go build -o math

CMD ["./math"]