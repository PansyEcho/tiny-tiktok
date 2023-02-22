package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"time"
	"tiny-tiktok/common/cos"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/video/internal/model"
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

func (l *PubVideoLogic) PubVideo() (resp *types.PublishVideoResp, err error) {

	token := l.r.FormValue("token")
	if token == "" {
		return nil, errors.New(errx.MapErrMsg(errx.NOT_TOKEN_ERROR))
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return nil, errors.New(errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR))
	}

	file, _, err := l.r.FormFile("data")

	upLoader := cos.UploaderVideo{
		UserId:      userClaim.Id,
		MachineId:   l.svcCtx.Config.COSConf.MachineId,
		VideoBucket: l.svcCtx.Config.COSConf.VideoBucket,
		SecretID:    l.svcCtx.Config.COSConf.SecretId,
		SecretKey:   l.svcCtx.Config.COSConf.SecretKey,
	}

	key, err := upLoader.UploadVideo(l.ctx, file)
	if err != nil {
		return &types.PublishVideoResp{
			Status: types.Status{
				Status_code: errx.UPLOAD_VIDEO_ERROR,
				Status_msg:  errx.MapErrMsg(errx.UPLOAD_VIDEO_ERROR),
			},
		}, nil
	}

	video := &model.Video{
		AuthorID:      userClaim.Id,
		Title:         l.r.FormValue("title"),
		PlayURL:       l.svcCtx.Config.COSConf.VideoBucket + "/" + key + ".mp4",
		CoverURL:      l.svcCtx.Config.COSConf.CoverBucket + "/" + key + ".jpg",
		PublishTime:   time.Now(),
		FavoriteCount: 0,
		CommentCount:  0,
	}

	tx := l.svcCtx.DB.Create(&video)

	if tx.Error != nil || tx.RowsAffected == 0 {
		return &types.PublishVideoResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}

	tx = l.svcCtx.DB.Model(&model.User{}).Where("id = ?", userClaim.Id).UpdateColumn("work_count", gorm.Expr("work_count + ?", 1))
	if tx.Error != nil || tx.RowsAffected == 0 {
		return &types.PublishVideoResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}

	resp = new(types.PublishVideoResp)
	resp.Status_code = errx.SUCCEED
	resp.Status_msg = errx.MapErrMsg(errx.SUCCEED)
	return
}
