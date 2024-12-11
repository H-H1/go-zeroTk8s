package logic

import (
	"context"

	"tk8s/api/user/internal/svc"
	"tk8s/api/user/internal/types"
	"tk8s/gen"
	"tk8s/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserReply, err error) {
	// todo: add your logic here and delete this line
	UserResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	//测试下连接mysql
	var Testtable *gen.Testtable
	Testtable, err = l.svcCtx.T1111Model.FindOne(l.ctx, 1)

	if err != nil {
		return nil, err
	}
	return &types.UserReply{
		Id:   UserResp.Id,
		Name: Testtable.Name.String + "--" + UserResp.Name,
	}, err
}
