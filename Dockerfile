FROM alpine:3.19
LABEL maintainer="Marvin Menzerath <dev@marvin.pro>"

RUN apk add -U --no-cache ca-certificates

WORKDIR /app
COPY --chmod=0755 build/exporter .

ENTRYPOINT ["./exporter"]
