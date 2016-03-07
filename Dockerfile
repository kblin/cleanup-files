# A very minimal container that cleans up files
# VERSION 0.1.0
FROM scratch
MAINTAINER Kai Blin <kblin@biosustain.dtu.dk>

ADD ./cleanup-files /cleanup-files

VOLUME ["/data"]

ENTRYPOINT ["/cleanup-files"]
