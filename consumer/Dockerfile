FROM golang:1.19 AS builder
RUN apt update && apt install git && apt install tzdata
RUN apt install ca-certificates && update-ca-certificates
WORKDIR $GOPATH/src/pressus
COPY .. .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/pressus


FROM scratch

COPY --from=builder /go/bin/event_app_auth /go/bin/pressus
COPY --from=builder /go/src/pressus/config.yaml /config.yaml
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV TZ Europe/Moscow
CMD ["/go/bin/pressus"]