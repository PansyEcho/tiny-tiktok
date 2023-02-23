package relation

import (
	"context"
	"github.com/jinzhu/copier"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/relation/rpc/relationservice"

	"tiny-tiktok/service/relation/internal/svc"
	"tiny-tiktok/service/relation/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// todo: add your logic here and delete this line
	if req.Token == "" {
		actionCode(errx.UNLOGIN_ERROR, nil, err)
		return
	}
	userClaim, err := utils.AnalyzeToken(req.Token)
	if err != nil {
		actionCode(errx.TOKEN_EXPIRE_ERROR, nil, err)
		return
	}
	friendlist_rpc, err := l.svcCtx.RelationRpc.FriendList(l.ctx, &relationservice.FriendListReq{
		UserId: userClaim.Id,
	})
	var list []types.User
	for _, v := range friendlist_rpc.UserList {
		var item types.User
		_ = copier.Copy(&item, v)
		list = append(list, item)
	}
	if err != nil {
		actionCode(errx.RPC_ERROR, nil, err)
		return
	}
	return &types.FriendListResp{
		Status: types.Status{
			Status_code: friendlist_rpc.StatusCode,
			Status_msg:  friendlist_rpc.StatusMsg,
		},
		FriendList: list,
	}, nil
}
