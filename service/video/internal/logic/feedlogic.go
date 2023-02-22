package logic

import (
	"context"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/video/internal/model"
	"tiny-tiktok/service/video/internal/svc"
	"tiny-tiktok/service/video/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {
	token := req.Token
	userClaim := new(utils.UserClaim)
	if token == "" {
		userClaim = &utils.UserClaim{
			Id:       -11,
			Username: "",
		}
	}
	userClaim, err = utils.AnalyzeToken(token)
	if err != nil {
		userClaim = &utils.UserClaim{
			Id:       -11,
			Username: "",
		}
	}
	var videos []*model.Video

	find := l.svcCtx.DB.Limit(30).Find(&videos)

	if find.Error != nil || find.RowsAffected == 0 {
		return &types.FeedResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}

	var arrVideos []*types.Video

	for _, video := range videos {
		user := new(model.User)
		tx := l.svcCtx.DB.Where("id = ?", video.AuthorID).Limit(1).Find(&user)
		if tx.Error != nil || tx.RowsAffected == 0 {
			return &types.FeedResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}
		follow := new(model.Follow)

		tx = l.svcCtx.DB.Where("user_id = ? and follow_id = ?", userClaim.Id, user.ID).Limit(1).Find(&follow)
		if tx.Error != nil {
			return &types.FeedResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}
		favorite := new(model.Favorite)
		fa := l.svcCtx.DB.Where("user_id = ? and video_id = ? and cancel = 0", userClaim.Id, video.ID).Limit(1).Find(&favorite)
		if fa.Error != nil {
			return &types.FeedResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}

		videoTemp := &types.Video{
			Id:            video.ID,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    fa.RowsAffected == 1,
			Title:         video.Title,
			User: types.User{
				UserID:          user.ID,
				Username:        user.Username,
				FollowCount:     user.Follow_count,
				FollowerCount:   user.Follower_count,
				IsFollow:        tx.RowsAffected == 1,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				TotalFavorited:  user.TotalFavorited,
				WorkCount:       user.WorkCount,
				FavoriteCount:   user.FavoriteCount,
			},
		}
		arrVideos = append(arrVideos, videoTemp)

	}

	return &types.FeedResp{
		Status: types.Status{
			Status_code: errx.SUCCEED,
			Status_msg:  errx.MapErrMsg(errx.SUCCEED),
		},
		NextTime: 10,
		Video:    arrVideos,
	}, nil
}
