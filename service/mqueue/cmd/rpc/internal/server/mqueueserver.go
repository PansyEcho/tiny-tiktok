// Code generated by goctl. DO NOT EDIT.
// Source: mqueue.proto

package server

import (
	"context"

	"tiny-tiktok/service/mqueue/cmd/rpc/internal/logic"
	"tiny-tiktok/service/mqueue/cmd/rpc/internal/svc"
	"tiny-tiktok/service/mqueue/cmd/rpc/pb"
)

type MqueueServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMqueueServer
}

func NewMqueueServer(svcCtx *svc.ServiceContext) *MqueueServer {
	return &MqueueServer{
		svcCtx: svcCtx,
	}
}

func (s *MqueueServer) KqRelationAction(ctx context.Context, in *pb.KqRelationActionReq) (*pb.KqRelationActionResp, error) {
	l := logic.NewKqRelationActionLogic(ctx, s.svcCtx)
	return l.KqRelationAction(in)
}
