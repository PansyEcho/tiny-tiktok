package logic

import (
	"context"
	"github.com/jinzhu/copier"
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
	user := new(model.User)
	find := l.svcCtx.DB.Where("id = ?", userClaim.Id).Limit(1).Find(&user)
	if find.RowsAffected == 0 {
		resp.Status = types.Status{Status_code: errx.NOT_USER_ERROR, Status_msg: errx.MapErrMsg(errx.NOT_USER_ERROR)}
		return
	}
	follow := new(model.Follow)
	find = l.svcCtx.DB.Where("user_id = ? and follow_id = ? ").Limit(1).Find(&follow)
	userResp := new(types.User)
	err = copier.Copy(userResp, user)
	if err != nil {
		userinfoCode(errx.ERROR, resp, err)
		return
	}

	if find.RowsAffected > 0 && follow.Cancel == 1 {
		userResp.IsFollow = false
	}

	succeedUserInfo(userResp, resp)
	return
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
