FROM golang:1.20.3-bullseye as build

WORKDIR /

RUN set -eux; \
	apt-get update && apt-get upgrade && \
    	apt-get install -y git make cmake && \
	git clone https://github.com/kelseyhightower/confd.git && \
	cd confd && mkdir -p bin && make

FROM debian:bullseye

WORKDIR /

COPY --from=build /confd/bin/confd /bin/confd

RUN set -eux; \
	mkdir -p /etc/confd/conf.d

ENTRYPOINT [ "/bin/confd" ]