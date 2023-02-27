FROM golang:1.20.0

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o app-exe
