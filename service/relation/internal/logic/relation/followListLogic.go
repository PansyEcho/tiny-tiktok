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

type FollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowListLogic) FollowList(req *types.FollowListReq) (resp *types.FollowListResp, err error) {
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
	followlist_rpc, err := l.svcCtx.RelationRpc.FollowList(l.ctx, &relationservice.FollowListReq{
		UserId: userClaim.Id,
	})
	if err != nil {
		actionCode(errx.RPC_ERROR, nil, err)
		return
	}
	var list []types.User
	for _, v := range followlist_rpc.UserList {
		var item types.User
		_ = copier.Copy(&item, v)
		list = append(list, item)
	}
	return &types.FollowListResp{
		Status: types.Status{
			Status_code: followlist_rpc.StatusCode,
			Status_msg:  followlist_rpc.StatusMsg,
		},
		FollowList: list,
	}, nil
}
