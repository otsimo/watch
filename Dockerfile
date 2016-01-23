FROM centurylink/ca-certs
MAINTAINER Sercan Degirmenci <sercan@otsimo.com>

ADD bin/watch-linux-amd64 /opt/otsimo-watch/bin/watch

EXPOSE 18858
# enable verbose debug for now
CMD ["/opt/otsimo-watch/bin/watch","--debug"]
