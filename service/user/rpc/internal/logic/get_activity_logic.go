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

type GetActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityLogic {
	return &GetActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityLogic) GetActivity(in *user.GetActivityReq) (*user.GetActivityResp, error) {
	// todo: add your logic here and delete this line
	token := in.GetToken()
	if token == "" {
		return nil, errors.New(errx.MapErrMsg(errx.NOT_TOKEN_ERROR))
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return nil, errors.New(errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR))
	}
	userInfo := new(model.User)

	find := l.svcCtx.DB.Where("id = ?", userClaim.Id).Limit(1).Find(&userInfo)
	if find.Error != nil || find.RowsAffected == 0 {
		return nil, errors.New(errx.MapErrMsg(errx.DB_ERROR))
	}

	resp := new(user.GetActivityResp)

	resp.Value = userInfo.Activity

	return resp, nil
}
