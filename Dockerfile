FROM golang:1.13-alpine AS builder

ENV GOPROXY="direct"

WORKDIR /go/src/github.com/mtpereira/fizzbuzz
COPY . .

RUN adduser -D -g '' go \
        && apk add --no-cache git ca-certificates tzdata \
        && update-ca-certificates
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
    go build -ldflags='-w -s -extldflags "-static"' \
                -o /go/bin/fizzbuzz ./cmd/fizzbuzz \
        && chmod 500 /go/bin/fizzbuzz

FROM scratch

COPY --from=builder /etc/passwd /etc/group /etc/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder --chown=go /go/bin/fizzbuzz /go/bin/fizzbuzz

USER go
ENTRYPOINT [ "/go/bin/fizzbuzz" ]
