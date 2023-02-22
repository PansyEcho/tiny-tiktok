package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/video/internal/model"

	"tiny-tiktok/service/video/internal/svc"
	"tiny-tiktok/service/video/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishVideoListLogic {
	return &GetPublishVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublishVideoListLogic) GetPublishVideoList(req *types.GetPubVideoReq) (resp *types.GetPubVideoResp, err error) {
	token := req.Token
	if token == "" {
		return nil, errors.New(errx.MapErrMsg(errx.NOT_TOKEN_ERROR))
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return nil, errors.New(errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR))
	}

	var videos []model.Video

	if req.UserId == 0 {
		req.UserId = userClaim.Id
	}

	find := l.svcCtx.DB.Where("author_id = ?", req.UserId).Find(&videos)

	if find.Error != nil {
		return &types.GetPubVideoResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}

	var respVideos []*types.Video
	for _, video := range videos {
		user := new(model.User)
		follow := new(model.Follow)
		findB := l.svcCtx.DB.Where("id = ?", video.AuthorID).Find(&user)
		findC := l.svcCtx.DB.Where("user_id = ? and follow_id = ? and cancel = 0", userClaim.Id, video.AuthorID).Find(&follow)
		if findC.Error != nil || findB.Error != nil {
			return &types.GetPubVideoResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}
		respVideo := &types.Video{
			Id:            video.ID,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			CommentCount:  video.CommentCount,
			FavoriteCount: video.FavoriteCount,
			Title:         video.Title,
			IsFavorite:    false,
			User: types.User{
				UserID:          user.ID,
				Username:        user.Username,
				FollowCount:     user.Follow_count,
				FollowerCount:   user.Follower_count,
				IsFollow:        findC.RowsAffected == 1,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				TotalFavorited:  user.TotalFavorited,
				WorkCount:       user.WorkCount,
				FavoriteCount:   user.FavoriteCount,
			},
		}
		respVideos = append(respVideos, respVideo)
	}

	return &types.GetPubVideoResp{
		Status: types.Status{
			Status_code: errx.SUCCEED,
			Status_msg:  errx.MapErrMsg(errx.SUCCEED),
		},
		VideoPubList: respVideos,
	}, nil
}
