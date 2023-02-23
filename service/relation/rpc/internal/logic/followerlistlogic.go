package logic

import (
	"context"

	"tiny-tiktok/service/relation/rpc/internal/svc"
	"tiny-tiktok/service/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowerListLogic) FollowerList(in *relation.FollowerListReq) (*relation.FollowerListResp, error) {
	// todo: add your logic here and delete this line
	var followerlist []interface{}
	l.svcCtx.DB.Where("follow_id = ?", in.UserId).Select("id").Find(&followerlist)
	for _, v := range followerlist {
		//等userrpc
	}
	return &relation.FollowerListResp{}, nil
}
