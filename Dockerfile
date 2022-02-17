FROM alpine:3.15
LABEL maintainer="Marvin Menzerath <dev@marvin.pro>"

RUN apk add -U --no-cache ca-certificates

WORKDIR /app
COPY build/exporter .

ENTRYPOINT ["./exporter"]
