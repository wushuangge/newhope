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

WORKDIR /go/src/newhope

ADD . /go/src/newhope
RUN CGO_ENABLED=0 go build -o newhope /go/src/newhope/app/main.go

FROM alpine:latest AS final

WORKDIR /app
COPY --from=builder /go/src/newhope/newhope /app/

ENTRYPOINT ["/app/newhope"]