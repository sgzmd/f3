FROM tetafro/golang-gcc:latest

ENV test=0
ARG telegram_api_key=nonkey
ARG telegram_chat_id=-1
ARG mariadb_host=localhost
ARG mariadb_port=3306
ARG mariadb_user=flibustier
ARG mariadb_database=flibustier
ARG mariadb_password=flibusta

ENV TELEGRAM_API_KEY=${telegram_api_key}
ENV TELEGRAM_CHAT_ID=${telegram_chat_id}
ENV MARIADB_HOST=${mariadb_host}
ENV MARIADB_PORT=${mariadb_port}
ENV MARIADB_USER=${mariadb_user}
ENV MARIADB_DATABASE=${mariadb_database}
ENV MARIADB_PASSWORD=${mariadb_password}

RUN echo "================$telegram_api_key=================="

RUN mkdir /app
RUN apk add sqlite gawk bash wget curl

ADD . /app

WORKDIR /app

RUN go build -o downloader cmd/downloader/main.go
RUN go build -o flibustier_server -tags fts5 flibuserver/server/*.go

FROM alpine:latest

# Expects /var/local/flibudata volume to be set

RUN mkdir /app
RUN apk add sqlite gawk bash curl mariadb-client python3 py3-pip wget

ADD . /app

WORKDIR /app

COPY --from=0 /app/scripts/docker_download.sh /app/docker_download.sh
COPY --from=0 /app/downloader /app/downloader
COPY --from=0 /app/flibustier_server /app/flibustier_server

RUN pip install --no-cache-dir --upgrade pip && pip install --no-cache-dir mysql-to-sqlite3 mysql-connector-python
RUN wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.13/grpc_health_probe-linux-amd64
RUN chmod +x ./grpc_health_probe-linux-amd64
RUN ls -lha ./grpc_health_probe-linux-amd64

RUN crontab downloader.cron

EXPOSE 9000

CMD /bin/sh /app/start.sh
