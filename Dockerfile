# Do not build directly. Use ./scripts/release.sh

FROM alpine

RUN apk update && \
    apk add ca-certificates wget && \
    update-ca-certificates

ADD ./bin/koki-short-drone-plugin /bin/short-drone-plugin

ARG SHORT_VERSION

RUN wget https://github.com/koki/short/releases/download/${SHORT_VERSION}/short_linux_amd64 -O /bin/short

RUN chmod a+x /bin/short

ENTRYPOINT ["short-drone-plugin"]
