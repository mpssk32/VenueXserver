FROM golang:1.25

WORKDIR /app

COPY go.mod .

RUN go mod tidy

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]