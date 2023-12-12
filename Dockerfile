# syntax=docker/dockerfile:1
FROM ubuntu:22.04

# install app dependencies
RUN apt-get update && apt-get install -y tzdata

COPY main /
COPY date-utc* /

# ENV PROD=true
EXPOSE 8999
CMD ./main
