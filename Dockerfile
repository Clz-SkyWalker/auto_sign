FROM alpine:latest
LABEL author="ClzSkywalker"

RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

RUN mkdir -p /app
RUN mkdir -p /app/config
WORKDIR /app

ADD ./bin/auto_sign /app

CMD ["./auto_sign"]

# FROM ubuntu:latest
# LABEL author="ClzSkywalker"
#
# RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai'>/etc/timezone
#
# RUN  sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
#      apt-get clean && \
#      apt-get update
#     
# RUN mkdir /app
# RUN mkdir /app/config
# WORKDIR /app
#
# ADD ./bin/auto_sign /app
#
# CMD ["./auto_sign"]
#
