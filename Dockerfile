FROM alpine:3.4
MAINTAINER JaysonGe <gyj3023@foxmail.com>

RUN apk add --update \
    && rm -rf /var/cache/apk/*

ADD  ./image /bin/
COPY ./go-hello-world  /bin/go-hello-world

EXPOSE 80
CMD ["go-hello-world"]

