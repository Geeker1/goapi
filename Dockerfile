FROM ubuntu:18.04

RUN apt update && apt upgrade

WORKDIR /usr/local/src

# COPY DB file
COPY db/clubs.json /usr/local/src

COPY ./main /usr/local/src
COPY views /usr/local/src/views

ENV DB_PATH=/usr/local/src/clubs.json

ENTRYPOINT [ "/usr/local/src/main" ]
