FROM golang:1.14-alpine as builder
WORKDIR /go/src/github.com/moov-io/wire
RUN apk add -U make
RUN adduser -D -g '' --shell /bin/false moov
COPY . .
ENV GO111MODULE=on
RUN go mod download
RUN make build
USER moov

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/wire/bin/server /bin/server
COPY --from=builder /etc/passwd /etc/passwd
USER moov
EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/server"]
