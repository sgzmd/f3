#FROM tetafro/golang-gcc:latest
FROM golang:1.20.1-alpine3.17

# Including the revision image was built from
ARG vcs-ref=0
ENV vcs-ref=$vcs-ref

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
