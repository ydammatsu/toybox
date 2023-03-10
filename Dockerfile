FROM golang:1.19

ENV PORT 8080

RUN apt-get update && apt-get install -y \
    git

WORKDIR /app
COPY . .

RUN go install github.com/cosmtrek/air@latest

RUN go mod download
