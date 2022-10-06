FROM tetafro/golang-gcc:latest

ENV test=0
ENV telegram_api_key=nonkey
ENV telegram_chat_id=-1
ENV mariadb_host=localhost
ENV mariadb_port=3306
ENV mariadb_user=flibustier
ENV mariadb_database=flibustier
ENV mariadb_password=flibusta

RUN mkdir /app
RUN apk add sqlite gawk bash wget curl

ADD . /app

WORKDIR /app

RUN go build -o downloader cmd/downloader/main.go
RUN go build -o flibustier_server -tags fts5 flibuserver/server/*.go

FROM alpine:latest

# Expects /var/local/flibudata volume to be set

RUN mkdir /app
RUN apk add sqlite gawk bash curl mariadb-client python3 py3-pip

ADD . /app

WORKDIR /app

COPY --from=0 /app/scripts/docker_download.sh /app/docker_download.sh
COPY --from=0 /app/downloader /app/downloader
COPY --from=0 /app/flibustier_server /app/flibustier_server

RUN pip install --no-cache-dir --upgrade pip && pip install --no-cache-dir mysql-to-sqlite3

RUN crontab downloader.cron

EXPOSE 9000

CMD /bin/sh /app/start.sh
