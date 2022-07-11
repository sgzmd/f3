FROM tetafro/golang-gcc:latest

ENV telegram_api_key=nonkey

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o telegram_bot tg/telegram.go