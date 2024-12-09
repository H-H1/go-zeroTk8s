# 第一阶段：创建用户和组
FROM alpine:latest as builder

# 创建 prometheus 用户和组
RUN apk add --no-cache shadow && \
    groupadd -r prometheus && \
    useradd -r -g prometheus prometheus

# 第二阶段：使用 Prometheus 镜像
FROM prom/prometheus:v2.28.1

# 复制 Prometheus 配置文件
COPY /deploy/prometheus/server/prometheus_tk8s.yml /etc/prometheus/prometheus.yml
COPY /data/prometheus/data /prometheus

# 设置时区
ENV TZ=Asia/Shanghai

# 更改 /prometheus 目录的所有权
COPY --from=builder /etc/passwd /etc/group /etc/

# 临时切换到 root 用户
USER root

# 更改 /prometheus 目录的所有权
RUN chown -R prometheus:prometheus /prometheus

# 切换到 prometheus 用户
USER prometheus

# 指定 Prometheus 配置文件路径
CMD ["--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/prometheus"]