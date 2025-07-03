FROM ubuntu:24.04
LABEL org.opencontainers.image.authors="Jerrico Gamis <jecklgamis@gmail.com>"

RUN apt-get update -y

ENV APP_ENV=dev

RUN mkdir -p /app/bin
RUN mkdir -p /app/configs

COPY bin/server-linux-amd64 /app/bin/server
RUN  chmod +x /app/bin/*

COPY configs /app/configs
COPY server.key /app
COPY server.crt /app

WORKDIR /app
EXPOSE 8080
EXPOSE 8443

COPY docker-entrypoint.sh /
CMD ["/docker-entrypoint.sh"]

