FROM golang:1.15-buster as builder

COPY . .

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

RUN go build -o /go/bin/whatsmyip

FROM scratch

COPY --from=builder /go/bin/whatsmyip /whatsmyip

CMD [ "/whatsmyip" ]
