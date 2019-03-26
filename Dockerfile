FROM golang:1.11.4
COPY . /go/src/gitlab.com/team-monitoring/comedian
# Install dependencies
WORKDIR /go/src/gitlab.com/team-monitoring/comedian
# Compile comedian
RUN make build_linux

FROM debian:9.8
LABEL maintainer="Anatoliy Fedorenko <fedorenko.tolik@gmail.com>"
RUN  apt-get update \
  && apt-get install -y --no-install-recommends ca-certificates locales wget \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
COPY active.en.toml  /
COPY active.ru.toml  /  
COPY --from=0  /go/src/gitlab.com/team-monitoring/comedian/comedian /
COPY goose /
COPY migrations /migrations
COPY entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]
