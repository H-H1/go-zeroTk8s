# tk8s
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

