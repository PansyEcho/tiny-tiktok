package logic

import (
	"context"
	"tiny-tiktok/common/kqueue"

	"tiny-tiktok/service/mqueue/cmd/rpc/internal/svc"
	"tiny-tiktok/service/mqueue/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type KqRelationActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKqRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KqRelationActionLogic {
	return &KqRelationActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KqRelationActionLogic) KqRelationAction(in *pb.KqRelationActionReq) (*pb.KqRelationActionResp, error) {
	// todo: add your logic here and delete this line
	m := kqueue.ThirdPaymentUpdatePayStatusNotifyMessage{
		OrderSn:   in.Sn,
		PayStatus: in.Status,
	}

	if err := l.svcCtx.KqueueClient.Push(kqueue.Action_STATUS, m); err != nil {
		return nil, err
	}
	return &pb.KqRelationActionResp{}, nil
}
