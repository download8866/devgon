# 使用Go官方镜像作为构建环境
# FROM golang:1.22 as builder

# 设置工作目录
# WORKDIR /app

# 复制go.mod和go.sum文件并安装依赖
# COPY go.mod go.sum ./
# RUN go mod download

# 复制所有代码并编译
# COPY . ./
# RUN go build -o main .

# 创建一个较小的运行时镜像
# FROM alpine:latest

# WORKDIR /root/

# 将编译好的二进制文件复制到运行时镜像
# COPY --from=builder /app/main .

# 运行Gin应用
# CMD ["go", "run", "main.go"]
# CMD ["./main"]
FROM golang:latest
WORKDIR /app
