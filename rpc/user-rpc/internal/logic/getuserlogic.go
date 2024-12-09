package logic

import (
	"context"
	"os"

	"tk8s/rpc/user-rpc/internal/svc"
	"tk8s/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	// 调用user-rpc服务，测试k8S服务是否正常启动"
	return &user.UserResponse{Id: hostname, Name: hostname}, nil
}
