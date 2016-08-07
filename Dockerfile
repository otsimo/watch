FROM centurylink/ca-certs
MAINTAINER Sercan Degirmenci <sercan@otsimo.com>

ADD watch-linux-amd64 /opt/otsimo/watch

EXPOSE 18858

CMD ["/opt/otsimo/watch"]
