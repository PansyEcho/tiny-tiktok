package logic

import (
	"context"
	"tiny-tiktok/common/errx"
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
	// todo: add your logic here and delete this line
	resp = new(types.FeedResp)
	resp.Status = types.Status{Status_code: errx.SUCCEED, Status_msg: errx.MapErrMsg(errx.SUCCEED)}
	resp.NextTime = 1
	user := types.User{UserID: 1, Username: "shizhen", FollowCount: 1, FollowerCount: 1, IsFollow: false}
	//video := &types.Video{User: user, PlayURL: "https://v26-web.douyinvod.com/7aa9246f06bfe049c12e9e81d4477198/63e008f8/video/tos/cn/tos-cn-ve-15c001-alinc2/ow0fOAFsjrAMyI35By8DgknaAvA9tfJRhsbwCn/?a=6383&ch=5&cr=3&dr=0&lr=all&cd=0%7C0%7C0%7C3&cv=1&br=1528&bt=1528&cs=0&ds=4&ft=LjhJEL998xbtu4CmD0P5H4eaciDXt_IH85QEeWnW9mPD1Ini&mime_type=video_mp4&qs=0&rc=ZmUzNjUzNWloaDxpNGQ6O0BpajVxbWY6Zjk6aDMzNGkzM0AtXmEzX2M1X2ExMS9jY2EvYSMybDZpcjRnZnNgLS1kLS9zcw%3D%3D&l=202302060252179E2940547C1E1240A668&btag=8000", CoverURL: "", FavoriteCount: 1, CommentCount: 1, IsFavorite: true, Title: "标题A"}

	video := new(model.Video)
	videoResp := new(types.Video)

	l.svcCtx.DB.Where("id = 1").Find(&video)
	videoResp = &types.Video{
		Id:            video.ID,
		User:          user,
		PlayURL:       video.PlayURL,
		CoverURL:      video.CoverURL,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    true,
		Title:         video.Title,
	}

	arrVideos := make([]*types.Video, 1)
	arrVideos[0] = videoResp
	resp.Video = arrVideos

	return
}
