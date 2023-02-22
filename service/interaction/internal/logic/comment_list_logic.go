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

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	token := req.Token
	if token == "" {
		return &types.CommentListResp{
			Status: types.Status{
				Status_code: errx.NOT_TOKEN_ERROR,
				Status_msg:  errx.MapErrMsg(errx.NOT_TOKEN_ERROR),
			},
		}, nil
	}
	userClaim, err := utils.AnalyzeToken(token)
	if err != nil {
		return &types.CommentListResp{
			Status: types.Status{
				Status_code: errx.TOKEN_EXPIRE_ERROR,
				Status_msg:  errx.MapErrMsg(errx.TOKEN_EXPIRE_ERROR),
			},
		}, nil
	}

	var comments []*model.Comment
	tx := l.svcCtx.DB.Where("video_id = ?", req.VideoId).Order("created_at desc").Find(&comments)
	if tx.Error != nil {
		return &types.CommentListResp{
			Status: types.Status{
				Status_code: errx.DB_ERROR,
				Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
			},
		}, nil
	}
	var respCommentList []*types.Comment
	for _, comment := range comments {

		user := new(model.User)
		tx = l.svcCtx.DB.Where("id = ?", comment.UserID).Limit(1).Find(&user)
		if tx.Error != nil || tx.RowsAffected == 0 {
			return &types.CommentListResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}
		follow := new(model.Follow)
		tx = l.svcCtx.DB.Where("user_id = ? and follow_id = ?", userClaim.Id, user.ID).Limit(1).Find(&follow)
		if tx.Error != nil {
			return &types.CommentListResp{
				Status: types.Status{
					Status_code: errx.DB_ERROR,
					Status_msg:  errx.MapErrMsg(errx.DB_ERROR),
				},
			}, nil
		}

		respComment := &types.Comment{
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
		}
		respCommentList = append(respCommentList, respComment)
	}

	return &types.CommentListResp{
		Status: types.Status{
			Status_code: errx.SUCCEED,
			Status_msg:  errx.MapErrMsg(errx.SUCCEED),
		},
		CommentList: respCommentList,
	}, nil
}
