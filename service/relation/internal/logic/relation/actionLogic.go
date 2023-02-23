package relation

import (
	"context"
	"github.com/pkg/errors"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/relation/rpc/relationservice"

	"tiny-tiktok/service/relation/internal/svc"
	"tiny-tiktok/service/relation/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActionLogic) Action(req *types.ActionReq) (resp *types.ActionResp, err error) {
	// todo: add your logic here and delete this line
	if req.Token == "" {
		actionCode(errx.UNLOGIN_ERROR, resp, err)
		return
	}
	userClaim, err := utils.AnalyzeToken(req.Token)
	if err != nil {
		actionCode(errx.TOKEN_EXPIRE_ERROR, resp, err)
		return
	}
	relationrpc_resp, err := l.svcCtx.RelationRpc.Action(l.ctx, &relationservice.ActionReq{
		Fromid:     userClaim.Id,
		Followid:   req.To_user_id,
		Actiontype: req.Action_type,
	})
	if err != nil {
		actionCode(errx.RPC_ERROR, nil, err)
		return
	}
	return &types.ActionResp{
		types.Status{
			Status_code: relationrpc_resp.StatusCode,
			Status_msg:  relationrpc_resp.StatusMsg,
		},
	}, nil
}
func actionCode(code int32, resp *types.ActionResp, err error) {
	msg := errx.MapErrMsg(code)
	logx.Error(msg, err)
	err = errors.New(msg)
	resp.Status = types.Status{Status_msg: msg, Status_code: code}
}
