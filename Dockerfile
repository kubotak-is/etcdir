FROM alpine:3.7

ENV HOST="0.0.0.0"
ENV PORT="8080"

RUN apk add --no-cache ca-certificates

RUN wget https://github.com/evildecay/etcdkeeper/releases/download/v0.7.3/etcdkeeper-v0.7.3-linux_x86_64.zip
RUN unzip etcdkeeper-v0.7.3-linux_x86_64.zip
RUN chmod 755 /etcdkeeper/etcdkeeper

EXPOSE ${PORT}

WORKDIR /etcdkeeper

ENTRYPOINT ./etcdkeeper -h $HOST -p $PORT
