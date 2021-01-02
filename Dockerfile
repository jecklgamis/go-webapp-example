FROM ubuntu:18.04
MAINTAINER Jerrico Gamis <jecklgamis@gmail.com>

RUN apt-get update -y

ENV APP_ENVIRONMENT dev

EXPOSE 8080
EXPOSE 8443

RUN mkdir -p /app/bin
RUN mkdir -p /app/configs

COPY bin/server-linux-amd64 /app/bin/server
RUN  chmod +x /app/bin/*

COPY configs /app/configs
COPY server.key /app
COPY server.crt /app

WORKDIR /app
COPY docker-entrypoint.sh /
CMD ["/docker-entrypoint.sh"]

