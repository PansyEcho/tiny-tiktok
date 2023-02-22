package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/user/internal/model"

	"tiny-tiktok/service/user/internal/svc"
	"tiny-tiktok/service/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	resp = new(types.UserInfoResp)
	if req.Token == "" {
		userinfoCode(errx.UNLOGIN_ERROR, resp, err)
		return
	}

	userClaim, err := utils.AnalyzeToken(req.Token)
	if err != nil {
		userinfoCode(errx.TOKEN_EXPIRE_ERROR, resp, err)
		return
	}

	if req.UserID == 0 {
		req.UserID = userClaim.Id
	}

	user := new(model.User)
	find := l.svcCtx.DB.Where("id = ?", req.UserID).Limit(1).Find(&user)
	if find.Error != nil || find.RowsAffected == 0 {
		return &types.UserInfoResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}

	follow := new(model.Follow)
	find = l.svcCtx.DB.Where("user_id = ? and follow_id = ? ", userClaim.Id, user.ID).Limit(1).Find(&follow)
	if find.Error != nil {
		return &types.UserInfoResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}
	return &types.UserInfoResp{
		Status: types.Status{
			Status_code: errx.SUCCEED,
			Status_msg:  errx.MapErrMsg(errx.SUCCEED),
		},
		User: &types.User{
			UserID:          user.ID,
			Username:        user.Username,
			FollowCount:     user.Follow_count,
			FollowerCount:   user.Follower_count,
			IsFollow:        find.RowsAffected == 1,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		},
	}, nil
}

func succeedUserInfo(user *types.User, resp *types.UserInfoResp) {
	resp.Status = types.Status{Status_code: errx.SUCCEED, Status_msg: errx.MapErrMsg(errx.SUCCEED)}
	resp.User = user
}

func userinfoCode(code int32, resp *types.UserInfoResp, err error) {
	msg := errx.MapErrMsg(code)
	logx.Error(msg, err)
	err = errors.New(msg)
	resp.Status = types.Status{Status_msg: msg, Status_code: code}
}
