FROM alpine

RUN apk --update add openssl ca-certificates

COPY bin/random /srv/random

CMD /srv/random
