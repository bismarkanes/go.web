# set the base image to Debian
# https://hub.docker.com/_/debian/
FROM debian:stretch

# replace shell with bash so we can source files
RUN rm /bin/sh && ln -s /bin/bash /bin/sh

# update the repository sources list
# and install dependencies
RUN apt-get update \
&& DEBIAN_FRONTEND=noninteractive apt-get upgrade -y \
&& apt-get -y autoclean

ENV APP_NAME go.web
ENV APP_DIR /var/www

COPY go.web $APP_DIR/

WORKDIR $APP_DIR

EXPOSE 8080

ENTRYPOINT $APP_DIR/$APP_NAME
