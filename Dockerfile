FROM golang:1.19-alpine AS builder
COPY . /src
WORKDIR /src
RUN mkdir -p bin/ && go build -o ./bin/ ./...

FROM alpine
LABEL name=unifi-ipv6-firewall-updater
LABEL url=https://github.com/suncj9807/unifi-ipv6-firewall-updater
COPY --from=builder /src/bin /app
COPY config.yaml /conf/config.yaml
ENV CONFIG_FILE_PATH /conf/config.yaml
WORKDIR /app
EXPOSE 28765
VOLUME /conf
ENTRYPOINT ["./unifi-ipv6-firewall-updater"]