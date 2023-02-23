FROM golang:1.19

RUN apt-get update && apt-get install -y \
    git

WORKDIR /app
COPY . .
