package logic

import (
	"context"
	"time"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/interaction/internal/model"

	"tiny-tiktok/service/interaction/internal/svc"
	"tiny-tiktok/service/interaction/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.CommentReq) (resp *types.CommentResp, err error) {
	token := req.Token
	if token == "" {
		return &types.CommentResp{
			Status: types.Status{
				Status_code: errx.NOT_TOKEN_ERROR,
				Status_msg:  errx.MapErrMsg(errx.NOT_TOKEN_ERROR),
			},
		}, nil
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return &types.CommentResp{
			Status: types.Status{
				Status_code: errx.TOKEN_EXPIRE_ERROR,
				Status_msg:  errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR),
			},
		}, nil
	}

	comment := &model.Comment{
		UserID:  userClaim.Id,
		VideoID: req.VideoId,
		Content: req.CommentText,
		Cancel:  1,
	}

	if req.ActionType == 2 {
		isPriDel := false
		if req.CommentId == userClaim.Id {
			isPriDel = true
		}
		if !isPriDel {
			find := l.svcCtx.DB.Model(&model.Video{}).Where("video_id = ? and author_id = ?", req.VideoId, userClaim.Id).Limit(1)
			if find.Error != nil {
				return &types.CommentResp{
					Status: types.Status{
						Status_code: errx.DB_ERROR,
						Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
					},
				}, nil
			}
			if find.RowsAffected != 0 {
				isPriDel = true
			}
		}

		if !isPriDel {
			return &types.CommentResp{
				Status: types.Status{
					Status_code: errx.NOT_PRIVILEGES_ERROR,
					Status_msg:  errx.MapErrMsg(errx.NOT_PRIVILEGES_ERROR),
				},
			}, nil
		}

		tx := l.svcCtx.DB.Model(&comment).Update("cancel", 1)
		if tx.Error != nil {
			return &types.CommentResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}
	} else {
		comment.CreateDate = time.Now().Format("01-02")
		tx := l.svcCtx.DB.Create(&comment)
		if tx.Error != nil {
			return &types.CommentResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}
	}

	user := new(model.User)
	tx := l.svcCtx.DB.Where("id = ?", userClaim.Id).Limit(1).Find(&user)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return &types.CommentResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}
	follow := new(model.Follow)
	tx = l.svcCtx.DB.Where("user_id = ? and follow_id = ?", userClaim.Id, user.ID).Limit(1).Find(&follow)
	if tx.Error != nil {
		return &types.CommentResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}

	return &types.CommentResp{
		Status: types.Status{
			Status_code: errx.SUCCEED,
			Status_msg:  errx.MapErrMsg(errx.SUCCEED),
		},
		Comment: &types.Comment{
			CommentId:  comment.ID,
			Content:    comment.Content,
			CreateTime: comment.CreateDate,
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
		},
	}, nil
}
