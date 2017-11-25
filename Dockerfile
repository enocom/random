FROM alpine:latest

RUN apk --update add openssl ca-certificates

COPY bin/random /srv/random

ENV PORT="8080"

ENV POLL_DURATION="1m"

CMD ["/srv/random"]
