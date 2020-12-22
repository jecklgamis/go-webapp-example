FROM ubuntu:18.04
MAINTAINER Jerrico Gamis <jecklgamis@gmail.com>

RUN apt-get update -y

EXPOSE 8080
RUN mkdir -p /app/bin
COPY bin/server-linux-amd64 /app/bin/server
RUN  chmod +x /app/bin/*

WORKDIR /app
COPY docker-entrypoint.sh /
CMD ["/docker-entrypoint.sh"]

