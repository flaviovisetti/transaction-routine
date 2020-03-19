FROM golang:1.14-buster
LABEL maintainer="Flávio Visetti Soares"

ENV SRC_DIR=/go/src/github.com/flaviovisetti/transaction-routine

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - \
  && echo "deb https://packagecloud.io/golang-migrate/migrate/debian/ buster main" > /etc/apt/sources.list.d/migrate.list \
  && apt-get update \
  && apt-get install -y migrate

WORKDIR $SRC_DIR
COPY . ./