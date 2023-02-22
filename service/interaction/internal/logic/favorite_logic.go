package logic

import (
	"context"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/interaction/internal/model"

	"tiny-tiktok/service/interaction/internal/svc"
	"tiny-tiktok/service/interaction/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteLogic) Favorite(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {

	token := req.Token
	if token == "" {
		return &types.FavoriteResp{
			Status: types.Status{
				Status_code: errx.NOT_TOKEN_ERROR,
				Status_msg:  errx.MapErrMsg(errx.NOT_TOKEN_ERROR),
			},
		}, nil
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return &types.FavoriteResp{
			Status: types.Status{
				Status_code: errx.TOKEN_EXPIRE_ERROR,
				Status_msg:  errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR),
			},
		}, nil
	}

	video := new(model.Video)

	tx := l.svcCtx.DB.Where("id = ?", req.VideoId).Find(&video)
	if tx.Error != nil {
		return &types.FavoriteResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}

	favorite := &model.Favorite{
		UserID:  userClaim.Id,
		VideoID: req.VideoId,
		Cancel:  0,
	}

	if req.ActionType == 2 {
		tx = l.svcCtx.DB.Model(&favorite).Update("cancel", 1)
		if tx.Error != nil || tx.RowsAffected == 0 {
			return &types.FavoriteResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}
	} else {
		tx = l.svcCtx.DB.Create(&favorite)
		if tx.Error != nil || tx.RowsAffected == 0 {
			return &types.FavoriteResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}
	}

	return &types.FavoriteResp{
		Status: types.Status{
			Status_code: errx.SUCCEED,
			Status_msg:  errx.MapErrMsg(errx.SUCCEED),
		},
	}, nil
}
