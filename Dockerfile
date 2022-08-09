# FROM alpine:latest
# alpine uses 'musl libc' rather than 'gnu libc', /lib64/ doesn't exist, it can be used by a soft link.
# apk mirror
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && ln -s /lib/libc.musl-x86_64.so.1 /lib/ld-linux-x86-64.so.2 && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories 

# RUN mkdir -p /app
# RUN mkdir -p /app/config
# WORKDIR /app
#
# ADD ./bin/auto_sign /app
#
# CMD ["nohup ./auto_sign > auto_sign.log 2>&1 &"]

FROM ubuntu:latest

RUN cp /etc/apt/sources.list /etc/apt/sources.list.bak \
    && echo "" > /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse" >> /etc/apt/sources.list \
    && apt-get update
    
RUN mkdir /app
RUN mkdir /app/config
WORKDIR /app

ADD ./bin/auto_sign /app

CMD ["./auto_sign"]

