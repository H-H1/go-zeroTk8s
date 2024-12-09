# 使用官方的 MySQL 8.0.28 镜像作为基础镜像
FROM mysql/mysql-server:8.0.28

# 设置时区为上海
ENV TZ=Asia/Shanghai

# 设置 MySQL root 用户的密码
ENV MYSQL_ROOT_PASSWORD=PXDN93VRKUm8TeE7

# 设置 MySQL 配置
CMD ["--default-authentication-plugin=mysql_native_password", \
     "--character-set-server=utf8mb4", \
     "--collation-server=utf8mb4_general_ci", \
     "--explicit_defaults_for_timestamp=true", \
     "--lower_case_table_names=1"]

# 挂载数据卷
VOLUME ["/var/lib/mysql"]

# 暴露 MySQL 端口
EXPOSE 3306

