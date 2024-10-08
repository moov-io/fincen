FROM golang:1.23-alpine as builder
WORKDIR /go/src/github.com/moov-io/fincen
RUN apk add -U git make
RUN adduser -D -g '' --shell /bin/false moov
COPY . .
RUN make build
USER moov

FROM scratch
LABEL maintainer="Moov <oss@moov.io>"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/fincen/bin/server /bin/server
COPY --from=builder /go/src/github.com/moov-io/fincen/configs/config.default.yml /configs/config.default.yml
COPY --from=builder /etc/passwd /etc/passwd

USER moov
EXPOSE 8206
EXPOSE 8207
ENTRYPOINT ["/bin/server"]
