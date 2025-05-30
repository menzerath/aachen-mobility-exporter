FROM alpine:3.22
LABEL maintainer="Marvin Menzerath <dev@marvin.pro>"

ENV MODE=production

RUN apk add -U --no-cache ca-certificates

WORKDIR /app
COPY --chmod=0755 build/exporter .

ENTRYPOINT ["./exporter"]
