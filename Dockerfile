FROM golang:latest

WORKDIR /WD

COPY go.mod go.sum ./

RUN go mod download

COPY . .