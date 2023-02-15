package logic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"

	"tiny-tiktok/service/video/internal/svc"
	"tiny-tiktok/service/video/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PubVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewPubVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *PubVideoLogic {
	return &PubVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *PubVideoLogic) PubVideo(req *types.PublishVideoReq) (resp *types.PublishVideoResp, err error) {

	file, _, err := l.r.FormFile("data")
	if err != nil {
		return nil, errors.New(errx.MapErrMsg(errx.WRONG_DATA_ERROR))
	}

	token := req.Token
	if token == "" {
		return nil, errors.New(errx.MapErrMsg(errx.NOT_TOKEN_ERROR))
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return nil, errors.New(errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR))
	}
	fmt.Println(file)
	fmt.Println(userClaim.Id)
	return
}
