FROM golang:1.20

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o "blog-service"

EXPOSE 50051

CMD ["air"]
