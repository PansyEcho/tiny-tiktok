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

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerListReq) (resp *types.FollowerListResp, err error) {
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
	followerlist_rpc, err := l.svcCtx.RelationRpc.FollowerList(l.ctx, &relationservice.FollowerListReq{
		UserId: userClaim.Id,
	})
	if err != nil {
		actionCode(errx.RPC_ERROR, nil, err)
		return
	}
	var list []types.User
	for _, v := range followerlist_rpc.UserList {
		var item types.User
		_ = copier.Copy(&item, v)
		list = append(list, item)
	}
	return &types.FollowerListResp{
		Status: types.Status{
			Status_code: followerlist_rpc.StatusCode,
			Status_msg:  followerlist_rpc.StatusMsg,
		},
		FollowerList: list,
	}, nil
}
