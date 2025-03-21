# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.24.1-alpine

# 设置工作目录
WORKDIR /app

# 将 Go 代码复制到容器中
COPY . .

# 构建 Go 应用
RUN go build -o auth .

# 暴露端口
EXPOSE 8082

# 运行 Go 应用
CMD ["./auth"]