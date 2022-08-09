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

RUN  sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
     apt-get clean && \
     apt-get update
    
RUN mkdir /app
RUN mkdir /app/config
WORKDIR /app

ADD ./bin/auto_sign /app

CMD ["./auto_sign"]

