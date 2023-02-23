package logic

import (
	"context"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/service/relation/rpc/internal/svc"
	"tiny-tiktok/service/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowListLogic) FollowList(in *relation.FollowListReq) (*relation.FollowListResp, error) {
	// todo: add your logic here and delete this line
	if num, err := l.svcCtx.RedisCache.Scard(string(in.UserId)); num == 0 {
		if err != nil {
			return &relation.FollowListResp{
				StatusCode: errx.REDIS_ERROR,
				StatusMsg:  errx.MapErrMsg(errx.REDIS_ERROR),
			}, err
		}
		var followlist []interface{}
		l.svcCtx.DB.Where("id = ?", in.UserId).Select("follow_id").Find(&followlist)
		for _, k := range followlist {
			l.svcCtx.RedisCache.Sadd(string(in.UserId), k)
		}
	}
	tem, err := l.svcCtx.RedisCache.Smembers(string(in.UserId))
	if err != nil {
		return &relation.FollowListResp{
			StatusCode: errx.REDIS_ERROR,
			StatusMsg:  errx.MapErrMsg(errx.REDIS_ERROR),
		}, err
	}
	//var Userlist []*types.User
	for _, v := range tem {
		//等userrpc
	}
	return &relation.FollowListResp{}, nil
}
