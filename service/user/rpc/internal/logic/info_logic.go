package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/user/rpc/internal/model"

	"tiny-tiktok/service/user/rpc/internal/svc"
	"tiny-tiktok/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type userInfoReq user.UserInfoReq

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------user-----------------------
func (l *InfoLogic) Info(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	token := in.GetToken()
	if token == "" {
		return nil, errors.New(errx.MapErrMsg(errx.NOT_TOKEN_ERROR))
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return nil, errors.New(errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR))
	}

	userSearch := new(model.User)
	find := l.svcCtx.DB.Where("id = ?", in.UserId).Limit(1).Find(&userSearch)
	if find.RowsAffected <= 0 {
		return nil, errors.New(errx.MapErrMsg(errx.NOT_USER_ERROR))
	}

	follow := new(model.Follow)
	isFollow := false
	find = l.svcCtx.DB.Where("user_id = ? and follow_id = ? and cancel = 0", userClaim.Id, in.UserId).Limit(1).Find(&follow)
	if find.RowsAffected > 0 {
		isFollow = true
	}

	userResp := new(user.UserInfoResp)
	userEntity := &user.User{
		UserId:        userSearch.ID,
		UserName:      userSearch.Username,
		FollowCount:   userSearch.Follow_count,
		FollowerCount: userSearch.Follower_count,
		IsFollow:      isFollow,
	}
	userResp.User = userEntity
	println(userResp.User.String())
	return userResp, nil

}

func (x *userInfoReq) IsValidToken() (bool, *utils.UserClaim, error) {
	token := x.Token
	if token == "" {
		return false, nil, errors.New(errx.MapErrMsg(errx.NOT_TOKEN_ERROR))
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return false, nil, errors.New(errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR))
	}
	return true, userClaim, nil
}
