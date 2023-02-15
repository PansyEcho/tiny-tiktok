package logic

import (
	"context"
	"fmt"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/service/user/rpc/types/user"

	"tiny-tiktok/service/interaction/internal/svc"
	"tiny-tiktok/service/interaction/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InteractionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInteractionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InteractionLogic {
	return &InteractionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InteractionLogic) Interaction(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIsInVzZXJuYW1lIjoiaHVhbmdkYW4iLCJleHAiOjE2NzY3MDc3OTZ9.z6TP8X2AxkobupgNTEv0k7-pifUwNkGjOHiiVN41h18"
	userInfoReq := &user.UserInfoReq{UserId: 1, Token: token}
	info, err := l.svcCtx.UserRpc.Info(l.ctx, userInfoReq)
	if err != nil {
		logx.Error("rpc错误")
	}

	fmt.Println("用户信息：" + info.User.String())

	addActivity := &user.UpdateActivityReq{Token: token, Value: 2}

	activity, err := l.svcCtx.UserRpc.AddActivity(l.ctx, addActivity)
	if err != nil || activity.Code != int64(errx.SUCCEED) {
		logx.Error("rpc错误")
	}

	getActivity := &user.GetActivityReq{Token: token}
	activityResp, err := l.svcCtx.UserRpc.GetActivity(l.ctx, getActivity)
	if err != nil || activity.Code != int64(errx.SUCCEED) {
		logx.Error("rpc错误")
	}

	fmt.Printf("用户活跃度为: %d\n", activityResp.Value)

	return
}
