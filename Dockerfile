FROM golang:1.22-alpine

WORKDIR /app

# 代理加速Go Module下载
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]