8# tk8s
 go-zero-k8s脚手架

- api/user/*.yaml根据自己的需求填写

  ```yaml
  Name: user
  Host: 0.0.0.0
  Port: 8887
  Mode: dev
  
  # RPC配置
  UserRpcConf:
    Endpoints:
      - 127.0.0.1:8886   # 本地
    NonBlock: True
    # Target: K8s://tk8s/base-svc:8886   #k8s模式下，指定目标服务
  
  #数据库、缓存配置
  DB:
    DataSource: root:PXDN93VRKUm8TeE7@tcp(127.0.0.1:33069)/t1111?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai  
    # DataSource: root:PXDN93VRKUm8TeE7@tcp(mysql-master:3306)/t1111?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
  ```

- 修改为本地环境测试跑下

  ```go
  api\user> go run user.go
  
  api\user> go run user.go
  
  测试下跑通
  ```

![image-20241210035953365](https://github.com/user-attachments/assets/f524ec80-f79a-40df-8de2-4296ff783765)





- 根据服务下单Dockerfile把api和Dockerfile 镜像打出来


![image-20241210041708249](https://github.com/user-attachments/assets/17b2e4cc-87d7-4855-94fd-5b8e678387bd)

1. 打好镜像后，k8ts执行项目下单api和rpc的yaml	
2. 以及go-zeroTk8s\deploy的的yaml，记得改api和rpc的*.yaml

会得到

![698285c3e397575075f2951fd1b8b10](https://github.com/user-attachments/assets/186c5c6d-5d26-4b7b-9b6f-6a8e719a37ba)


![image-20241210044704294](https://github.com/user-attachments/assets/0d019277-4e20-4ae5-9795-3d2621e5526a)



- 查看运行状态	

  可以暴露指定的k8s端口，也可以直接使用nodeport

```go
kubectl port-forward svc/base-api-svc 8887:8887 -n tk8s   #使用port-forward，

kubectl port-forward svc/prometheus-service 5556:9090 -n tk8s

```

测试下服务是否在k8s运行正常，可以看到三个不同的Hostname
![image-20241210045153866](https://github.com/user-attachments/assets/cee74bd6-ac72-4617-b4db-1b072be6f92a)
![70a17860af4ff41ef1d162574a2cef2](https://github.com/user-attachments/assets/d75417c1-a175-468e-9d42-764c217ffa43)
![8b5329bc9f599799c1e548c827a0ecb](https://github.com/user-attachments/assets/96d99380-b4da-457c-8849-f7e96b40d6bc)



浏览器打开，可以看到pmetheus启动

![7f788a3e49193492fe7add2747c8be0](https://github.com/user-attachments/assets/1581c4e4-5fa4-4b3d-8619-378a631f0b53)




## 主从复制配置

要在 Kubernetes 中配置 MySQL 的主从复制，你需要进行以下步骤：

1. **创建主节点和从节点的 StatefulSet**
2. **配置主节点和从节点的 MySQL 配置文件**
3. **初始化从节点的数据**

目前还需要完成3， 初始化从节点的数据

在主节点上创建一个用于复制的用户，并授权给从节点。

```sql
sh：
mysql -uroot -pPXDN93VRKUm8TeE7

sql 
查看是否生效my.cnf

SHOW VARIABLES LIKE 'server_id';
SHOW VARIABLES LIKE 'log_bin';
SHOW VARIABLES LIKE 'binlog_format';
SHOW VARIABLES LIKE 'gtid_mode';
SHOW VARIABLES LIKE 'enforce-gtid-consistency';

创建用于从服务器连接主服务器的用户

sql

CREATE USER 'root'@'%' IDENTIFIED BY 'PXDN93VRKUm8TeE7';
GRANT REPLICATION SLAVE ON *.* TO 'root'@'%';
FLUSH PRIVILEGES;

```

然后在主节点上获取二进制日志的位置信息。

```sql
sql

SHOW MASTER STATUS;
```

记录下 `File` 和 `Position` 的值，例如：

```sql
+------------------+----------+--------------+------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB |
+------------------+----------+--------------+------------------+
| binlog.000002    | 4007      |              |                  |
+------------------+----------+--------------+------------------+
```

在从节点上执行以下命令来启动复制：

```sql
sql

STOP SLAVE IO_THREAD FOR CHANNEL '';
CHANGE MASTER TO
  MASTER_HOST='mysql-master',
  MASTER_USER='root',
  MASTER_PASSWORD='PXDN93VRKUm8TeE7',
  MASTER_LOG_FILE='mysql-bin.000002',
  MASTER_LOG_POS=4007;
START SLAVE IO_THREAD FOR CHANNEL '';
START SLAVE;    #开始slave
SHOW SLAVE STATUS \G;  #查看详情，是否错误

```

导入数据，测试是否主从配置成功

```sql
#主节点
CREATE DATABASE t1111;
USE t1111;
CREATE TABLE testtable (id INT PRIMARY KEY, name VARCHAR(100));
INSERT INTO testtable VALUES (888, 'Test Data');


```

可以看到master上插入
![image-20241210051737759](https://github.com/user-attachments/assets/71acb111-02d2-469a-9e53-e649050b3c66)



slave上完成主从复制
![image-20241210051801418](https://github.com/user-attachments/assets/4c4177e8-170d-4190-bfde-b527ea12f24d)

