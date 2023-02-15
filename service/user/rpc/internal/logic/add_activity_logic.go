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

type AddActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityLogic {
	return &AddActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddActivityLogic) AddActivity(in *user.UpdateActivityReq) (*user.UpdateActivityResp, error) {
	userInfoResp := new(user.UpdateActivityResp)
	token := in.GetToken()
	if token == "" {
		return nil, errors.New(errx.MapErrMsg(errx.NOT_TOKEN_ERROR))
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return nil, errors.New(errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR))
	}
	userUpdate := new(model.User)
	find := l.svcCtx.DB.Where("id = ?", userClaim.Id).Limit(1).Find(&userUpdate)
	if find.Error != nil || find.RowsAffected == 0 {
		return nil, errors.New(errx.MapErrMsg(errx.NOT_USER_ERROR))
	}
	userUpdate.Activity += 1
	updates := l.svcCtx.DB.Model(&userUpdate).Where("id = ?", userUpdate.ID).Updates(userUpdate)
	if updates.Error != nil || updates.RowsAffected == 0 {
		return nil, errors.New(errx.MapErrMsg(errx.UPDATE_ACTIVITY_ERROR))
	}

	userInfoResp.Code = int64(errx.SUCCEED)
	userInfoResp.Msg = errx.MapErrMsg(errx.SUCCEED)

	return userInfoResp, nil
}
