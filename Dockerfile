FROM alpine:3.4
MAINTAINER Sercan Degirmenci <sercan@otsimo.com>

RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*


ADD watch-linux-amd64 /opt/otsimo/watch

EXPOSE 18858

CMD ["/opt/otsimo/watch"]
