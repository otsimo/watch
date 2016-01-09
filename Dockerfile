FROM centurylink/ca-certs
MAINTAINER Sercan Degirmenci <sercan@otsimo.com>

ADD bin/otsimo-watch-linux-amd64 /opt/otsimo-watch/bin/otsimo-awatchpi

# enable verbose debug for now
CMD ["/opt/otsimo-watch/bin/otsimo-watch","--debug"]
