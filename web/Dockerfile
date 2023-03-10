FROM golang:1.20.1-alpine3.17

ENV AUTH_MODE=""

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o web webserver/webserver.go

EXPOSE ${web_port}

RUN echo "================ web ${TELEGRAM_API_KEY} =================="

ENTRYPOINT /app/web ${AUTH_MODE} --grpc_backend=${BACKEND_URL} \
    --telegram_token=${TELEGRAM_API_KEY} \
    --web_port=${WEB_PORT} \
    --bot_name=${BOT_NAME} \
    --domain_name=${DOMAIN_NAME}