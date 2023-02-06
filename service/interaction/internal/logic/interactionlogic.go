package logic

import (
	"context"

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

	return
}
