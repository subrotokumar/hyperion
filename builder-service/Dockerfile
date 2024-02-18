FROM golang:alpine AS UPLOADER
WORKDIR /app
COPY main.go .
COPY go.mod .
COPY go.sum .
RUN go build -o main main.go

FROM ubuntu:focal

RUN apt-get update
RUN apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_20.x | bash -
RUN apt-get upgrade -y
RUN apt-get install -y nodejs
RUN apt-get install git -y

WORKDIR /home/app

COPY main.sh .
COPY --from=UPLOADER /app/main .
RUN chmod +x main.sh
RUN chmod +x main

ENTRYPOINT [ "/home/app/main.sh" ]