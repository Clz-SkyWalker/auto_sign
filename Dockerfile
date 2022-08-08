FROM alpine:latest

RUN mkdir -p /app
RUN mkdir -p /app/config
WORKDIR /app

ADD ./bin/auto_sign /app

CMD ["./auto_sign"]

