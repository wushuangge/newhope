FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/src/taskdash

ADD . /go/src/taskdash
RUN CGO_ENABLED=0 go build -o taskdash /go/src/taskdash/app/main.go

FROM alpine:latest AS final

WORKDIR /app
COPY --from=builder /go/src/taskdash/taskdash /app/
COPY --from=builder /go/src/taskdash/config /app/config
COPY --from=builder /go/src/taskdash/ui /app/ui

ENTRYPOINT ["/app/taskdash"]