FROM tetafro/golang-gcc:latest

ENV telegram_api_key=nonkey
ENV backend_url=http://localhost:9000
ENV web_port=8085
ENV bot_name=none
ENV domain_name=none

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o web webserver/webserver.go

FROM alpine:latest

RUN mkdir /app
ADD . /app

WORKDIR /app

COPY --from=0 /app/templates /app/templates
COPY --from=0 /app/web /app/web

EXPOSE 8080

ENTRYPOINT /app/web --grpc_backend=$backend_url --telegram_token=$telegram_api_key --web_port=$web_port --bot_name=$bot_name --domain_name=$domain_name