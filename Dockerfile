# 一阶段
FROM golang:1.25-alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.google.cn

RUN apk add --no-cache \
    gcc \
    musl-dev \
    git \
    make

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o rv .

# 二阶段
FROM alpine:latest AS runner

RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    && update-ca-certificates

RUN adduser -D -u 1000 appuser

WORKDIR /app

COPY --from=builder /build/rv .

RUN chown -R appuser:appuser /app

USER appuser

ENV TZ=Asia/Shanghai
ENV CGO_ENABLED=1

EXPOSE 8080 8081

# 入口点
ENTRYPOINT ["./rv"]
