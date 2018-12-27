########## Build ###################

FROM golang:1.11.2-alpine3.8 as builder
RUN apk update && apk add --no-cache git make gcc libc-dev

RUN adduser -D -g '' randomme

COPY . $GOPATH/src/github.com/richardcase/itsrandom/
WORKDIR $GOPATH/src/github.com/richardcase/itsrandom/

RUN make release


########## Output Image ###################
FROM scratch

COPY --from=builder /go/bin/itsrandom /app/itsrandom

USER randomme

EXPOSE 8080

ENTRYPOINT ["/app/itsrandom"]

