FROM alpine:latest
MAINTAINER Eno Compton <eno4@ecom.com>

RUN apk --update add openssl ca-certificates
COPY bin/random /srv/random
ENV ADDR=":8080"
ENV POLL_DURATION="1m"
CMD ["/srv/random"]
