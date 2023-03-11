#FROM tetafro/golang-gcc:latest
FROM golang:1.20.1-alpine3.17

# Including the revision image was built from
ARG vcs-ref=0
ENV vcs-ref=$vcs-ref

<<<<<<< HEAD
ARG telegram_api_key=nonkey
ARG telegram_chat_id=-1
ARG mariadb_host=localhost
ARG mariadb_port=3306
ARG mariadb_user=flibustier
ARG mariadb_database=flibustier
ARG mariadb_password=flibusta
ARG skip_build_flibustier=""
ARG flibusta_base_url="http://flibusta.site/sql"

ENV TELEGRAM_API_KEY=${telegram_api_key}
ENV TELEGRAM_CHAT_ID=${telegram_chat_id}
ENV MARIADB_HOST=${mariadb_host}
ENV MARIADB_PORT=${mariadb_port}
ENV MARIADB_USER=${mariadb_user}
ENV MARIADB_DATABASE=${mariadb_database}
ENV MARIADB_PASSWORD=${mariadb_password}
ENV FLIBUSTA_BASE_URL=${flibusta_base_url}

RUN echo "BASE URL: ${FLIBUSTA_BASE_URL}"

=======
>>>>>>> dev
RUN echo "================ ${TELEGRAM_API_KEY} =================="

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN ./build_flibuserver.sh


# Expects /var/local/flibudata volume to be set

RUN echo "================${TELEGRAM_API_KEY}=================="

RUN apk update
RUN apk upgrade

RUN apk add bash curl mariadb-client python3 wget

RUN wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.13/grpc_health_probe-linux-amd64
RUN chmod +x ./grpc_health_probe-linux-amd64
RUN ls -lha ./grpc_health_probe-linux-amd64

EXPOSE 9000

CMD /bin/sh /app/start.sh
