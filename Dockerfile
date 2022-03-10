FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

ENV GOARCH amd64

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o hello ./main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

ENV TZ Asia/Shanghai

WORKDIR /build

COPY --from=builder /build/hello /build/hello
COPY --from=builder /build/settings-dev.yaml /build/settings-dev.yaml
COPY --from=builder /build/logs /build/logs
COPY --from=builder /build/view /build/view
COPY --from=builder /build/public /build/public

CMD ["./hello","-f","./settings-dev.yaml"]