FROM golang:1.11.2-stretch
MAINTAINER Doni Rubiagatra <rubiagatra@gmail.com>

RUN mkdir -p /cloud-native-go-application
WORKDIR /cloud-native-go-application

COPY . .

EXPOSE 8080

RUN go mod tidy
ENTRYPOINT go run main.go

