########## Build ###################

FROM golang:1.11.2-alpine3.8 as builder
RUN apk update && apk add --no-cache git make gcc libc-dev

COPY . $GOPATH/src/github.com/richardcase/itsrandom/
WORKDIR $GOPATH/src/github.com/richardcase/itsrandom/

RUN make release


########## Output Image ###################
FROM scratch

COPY --from=builder /go/bin/itsrandom /app/itsrandom

EXPOSE 8080

ENTRYPOINT ["/app/itsrandom"]

