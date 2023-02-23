package logic

import (
	"context"

	"tiny-tiktok/service/relation/rpc/internal/svc"
	"tiny-tiktok/service/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *relation.FriendListReq) (*relation.FriendListResp, error) {
	// todo: add your logic here and delete this line
	var friendlist []interface{}
	l.svcCtx.DB.Where("follow_id = ?", in.UserId).Select("id").Find(&friendlist)
	for _, v := range friendlist {
		//等userrpc
	}
	return &relation.FriendListResp{}, nil
}
