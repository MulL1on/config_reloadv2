FROM golang:1.20.3-bullseye as build

WORKDIR /app

COPY . .

ENV GOPROXY=https://goproxy.cn

RUN mkdir /etc/confd

RUN go mod download

RUN go build app.go

CMD tail -f /dev/null

