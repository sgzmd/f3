FROM golang:1.20.1-alpine3.17

ARG telegram_api_key
ARG backend_url
ARG web_port
ARG bot_name
ARG domain_name

ENV TELEGRAM_API_KEY=${telegram_api_key}
ENV BACKEND_URL=${backend_url}
ENV WEB_PORT=${web_port}
ENV BOT_NAME=${bot_name}
ENV DOMAIN_NAME=${domain_name}

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o web webserver/webserver.go

EXPOSE ${web_port}

RUN echo "================ web ${TELEGRAM_API_KEY} =================="

ENTRYPOINT /app/web --grpc_backend=${BACKEND_URL} --telegram_token=${TELEGRAM_API_KEY} --web_port=${WEB_PORT} --bot_name=${BOT_NAME} --domain_name=${DOMAIN_NAME}