FROM golang:1.23-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 1
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc

RUN go build -tags musl -ldflags '-s -w -extldflags "-static"' -v -o /app/main ./main.go


FROM alpine

ARG VERSION
ENV VERSION ${VERSION}

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates bash tzdata curl
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/main /app/demo
COPY --from=builder /app/etc /app/etc

CMD ["./demo"]
