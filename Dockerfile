FROM golang:1.14-buster
LABEL maintainer="Flávio Visetti Soares"

ENV SRC_DIR=/go/src/github.com/flaviovisetti/transaction-routine

RUN apt-get update && apt-get install -y vim postgresql-client

WORKDIR $SRC_DIR
COPY . ./
RUN go get .