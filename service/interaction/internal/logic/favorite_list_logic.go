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

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListReq) (resp *types.FavoriteListResp, err error) {
	token := req.Token
	if token == "" {
		return &types.FavoriteListResp{
			Status: types.Status{
				Status_code: errx.NOT_TOKEN_ERROR,
				Status_msg:  errx.MapErrMsg(errx.NOT_TOKEN_ERROR),
			},
		}, nil
	}
	userCliam, err := utils.AnalyzeToken(token)
	if err != nil {
		return &types.FavoriteListResp{
			Status: types.Status{
				Status_code: errx.TOKEN_EXPIRE_ERROR,
				Status_msg:  errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR),
			},
		}, nil
	}

	var favorites []model.Favorite

	find := l.svcCtx.DB.Where("user_id = ? and cancel = 1", req.UserId).Find(&favorites)
	if find.Error != nil {
		return &types.FavoriteListResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}

	var respVideos []*types.Video

	for _, fa := range favorites {
		video := new(model.Video)
		user := new(model.User)
		follow := new(model.Follow)
		findA := l.svcCtx.DB.Where("id = ?", fa.VideoID).Find(&video)
		findB := l.svcCtx.DB.Where("id = ?", video.AuthorID).Find(&user)
		findC := l.svcCtx.DB.Where("user_id = ? and follow_id", userCliam.Id, video.AuthorID).Find(&follow)
		if findA.Error != nil || findB.Error != nil {
			return &types.FavoriteListResp{
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
			IsFavorite:    true,
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

	return &types.FavoriteListResp{
		Status: types.Status{
			Status_code: errx.SUCCEED,
			Status_msg:  errx.MapErrMsg(errx.SUCCEED),
		},
		FavoriteList: respVideos,
	}, nil
}
