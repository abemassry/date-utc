# syntax=docker/dockerfile:1
FROM ubuntu:22.04

# install app dependencies
RUN apt-get update && apt-get install -y tzdata
RUN mkdir /templates

COPY main /
COPY date-utc* /
COPY templates/index.tmpl /templates/index.tmpl

# ENV PROD=true
EXPOSE 80
CMD ./main
