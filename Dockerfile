FROM centurylink/ca-certs
MAINTAINER Sercan Degirmenci <sercan@otsimo.com>

ADD bin/watch-linux-amd64 /opt/otsimo-watch/bin/watch

# enable verbose debug for now
CMD ["/opt/otsimo-watch/bin/watch","--debug"]
