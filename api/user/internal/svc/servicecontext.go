package svc

import (
	"tk8s/api/user/internal/config"
	"tk8s/gen"
	"tk8s/rpc/user-rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userclient.User
	RedisClient *redis.Redis
	T1111Model  gen.TesttableModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	redisClient := redis.New(c.Cache.Host, func(r *redis.Redis) {
		r.Pass = c.Cache.Pass
	})
	return &ServiceContext{
		Config:      c,
		RedisClient: redisClient,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		T1111Model:  gen.NewTesttableModel(sqlConn, redisClient),
	}
}
