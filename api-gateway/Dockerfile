FROM golang:1.20

WORKDIR /app

# RUN go install github.com/cosmtrek/air@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o "/api-gateway"

# EXPOSE 8000

CMD ["/api-gateway"]
# CMD ["air"]
