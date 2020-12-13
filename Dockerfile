FROM golang:1.15-buster as builder

COPY . $GOPATH/src/cuddlydemos/whatsmyip/

WORKDIR $GOPATH/src/cuddlydemos/whatsmyip/

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

RUN go get -d -v
RUN go mod download && go build -o /go/bin/whatsmyip

FROM scratch

COPY --from=builder /go/bin/whatsmyip /whatsmyip

ENV GIN_MODE release
ENV PORT 8080

EXPOSE 8080

ENTRYPOINT [ "/whatsmyip" ]
