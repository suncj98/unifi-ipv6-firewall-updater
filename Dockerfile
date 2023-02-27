FROM golang:1.19 AS builder
COPY . /src
WORKDIR /src
RUN mkdir -p bin/ && go build -o ./bin/ ./...

FROM debian:stable-slim
COPY --from=builder /src/bin /app
WORKDIR /app
EXPOSE 8000
VOLUME /data/conf
CMD ["./cmd", "-conf", "/data/conf"]