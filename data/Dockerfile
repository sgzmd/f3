FROM tetafro/golang-gcc:latest

ENV test=0
ENV telegram_api_key=nonkey
ENV telegram_chat_id=-1

RUN mkdir /app
RUN apk add sqlite gawk bash wget curl

ADD . /app

WORKDIR /app

RUN go build -o downloader cmd/downloader/main.go
RUN go build -o flibustier_server -tags fts5 flibuserver/server/*.go

FROM alpine:latest

# Expects /var/local/flibudata volume to be set

RUN mkdir /app
RUN apk add sqlite gawk bash curl

ADD . /app

WORKDIR /app

COPY --from=0 /app/downloader /app/downloader
COPY --from=0 /app/flibustier_server /app/flibustier_server

RUN crontab downloader.cron

EXPOSE 9000

CMD /bin/sh /app/start.sh