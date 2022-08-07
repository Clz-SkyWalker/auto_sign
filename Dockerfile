FROM golang:1.18

MAINTAINER auto_sign
RUN mkdir -p /data/auto_sign
WORKDIR /data/auto_sign

EXPOSE 9000
CMD ["./bin/auto_sign"]


