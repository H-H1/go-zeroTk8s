# tk8s

English | [简体中文](README-cn.md)

go-zero-k8s Scaffold，more middleware will be added in the future



- [x] mysql
- [x] prometheu
- [ ] redis
- [ ] grafana
- [ ] ES
- [ ] jaeger
- [ ] kibana
- [ ] jaeger

- Fill in api/user/*.yaml according to your requirements.

  ```yaml
  Name: user
  Host: 0.0.0.0
  Port: 8887
  Mode: dev
  
  # RPC Configuration
  UserRpcConf:
    Endpoints:
      - 127.0.0.1:8886   # Local
    NonBlock: True
    # Target: K8s://tk8s/base-svc:8886   # In k8s mode, specify the target service
  
  # Database and Cache Configuration
  DB:
    DataSource: root:PXDN93VRKUm8TeE7@tcp(127.0.0.1:33069)/t1111?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai  
    # DataSource: root:PXDN93VRKUm8TeE7@tcp(mysql-master:3306)/t1111?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
  ```

- Modify for local environment testing and run.

  ```go
  api\user> go run user.go
  
  api\user> go run user.go
  
  ```

![image-20241210035953365](https://github.com/user-attachments/assets/f524ec80-f79a-40df-8de2-4296ff783765)





- Build the api and Dockerfile images according to the Dockerfile for service deployment.


![image-20241210041708249](https://github.com/user-attachments/assets/17b2e4cc-87d7-4855-94fd-5b8e678387bd)

1. After building the images, execute the api and rpc yaml files for the project in k8ts.
2. Also execute the yaml files in go-zeroTk8s\deploy. Remember to modify the api and rpc *.yaml files.
   You will get:



![698285c3e397575075f2951fd1b8b10](https://github.com/user-attachments/assets/186c5c6d-5d26-4b7b-9b6f-6a8e719a37ba)


![image-20241210044704294](https://github.com/user-attachments/assets/0d019277-4e20-4ae5-9795-3d2621e5526a)



- Check the running status.
You can expose the specified k8s port or directly use nodeport.

```go
kubectl port-forward svc/base-api-svc 8887:8887 -n tk8s   #使用port-forward，

kubectl port-forward svc/prometheus-service 5556:9090 -n tk8s

```

Test if the service is running normally in k8s. You can see three different Hostnames.
![image-20241210045153866](https://github.com/user-attachments/assets/cee74bd6-ac72-4617-b4db-1b072be6f92a)
![70a17860af4ff41ef1d162574a2cef2](https://github.com/user-attachments/assets/d75417c1-a175-468e-9d42-764c217ffa43)
![8b5329bc9f599799c1e548c827a0ecb](https://github.com/user-attachments/assets/96d99380-b4da-457c-8849-f7e96b40d6bc)



Open in the browser and you can see that pmetheus has started.

![7f788a3e49193492fe7add2747c8be0](https://github.com/user-attachments/assets/1581c4e4-5fa4-4b3d-8619-378a631f0b53)



## Master-Slave Replication Configuration



To configure MySQL master-slave replication in Kubernetes, you need to perform the following steps:



1. **Create StatefulSets for the master and slave nodes.**
2. **Configure the MySQL configuration files for the master and slave nodes.**
3. **Initialize the data for the slave node.**
   Currently, step 3, initializing the data for the slave node, still needs to be completed.
   Create a replication user on the master node and grant permissions to the slave node.

```sql
sh：
mysql -uroot -pPXDN93VRKUm8TeE7

sql 

SHOW VARIABLES LIKE 'server_id';
SHOW VARIABLES LIKE 'log_bin';
SHOW VARIABLES LIKE 'binlog_format';
SHOW VARIABLES LIKE 'gtid_mode';
SHOW VARIABLES LIKE 'enforce-gtid-consistency';



sql

CREATE USER 'root'@'%' IDENTIFIED BY 'PXDN93VRKUm8TeE7';
GRANT REPLICATION SLAVE ON *.* TO 'root'@'%';
FLUSH PRIVILEGES;

```

Then obtain the binary log position information on the master node.

```sql
sql

SHOW MASTER STATUS;
```

Record the values of `File` and `Position`, for example:

```sql
+------------------+----------+--------------+------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB |
+------------------+----------+--------------+------------------+
| binlog.000002    | 4007      |              |                  |
+------------------+----------+--------------+------------------+
```

Execute the following commands on the slave node to start replication:

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
START SLAVE;    # Start slave
SHOW SLAVE STATUS \G;  # Check details for errors

```

Import data and test if the master-slave configuration is successful.

```sql
# Master node
CREATE DATABASE t1111;
USE t1111;
CREATE TABLE testtable (id INT PRIMARY KEY, name VARCHAR(100));
INSERT INTO testtable VALUES (888, 'Test Data');


```

You can see the insertion on the master.
![image-20241210051737759](https://github.com/user-attachments/assets/71acb111-02d2-469a-9e53-e649050b3c66)



The slave completes the master-slave replication.
![image-20241210051801418](https://github.com/user-attachments/assets/4c4177e8-170d-4190-bfde-b527ea12f24d)

