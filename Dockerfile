# 第一阶段：构建阶段
FROM golang:1.23 as builder

WORKDIR /app

# 复制代码到容器中
COPY . .

# 编译应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 第二阶段：运行阶段
FROM alpine:latest

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .

# 设置容器启动命令
CMD ["/app/main"]