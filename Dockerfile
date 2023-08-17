FROM golang:1.19-alpine AS builder
COPY . /project
WORKDIR /project
RUN mkdir -p bin/ && go build -o ./bin/ ./...

FROM alpine
LABEL name=unifi-ipv6-firewall-updater
LABEL url=https://github.com/suncj98/unifi-ipv6-firewall-updater
COPY --from=builder /project/bin /app
ENV CONFIG_FILE_PATH /conf/config.yaml
WORKDIR /app
EXPOSE 28765
ENTRYPOINT ["./unifi-ipv6-firewall-updater"]