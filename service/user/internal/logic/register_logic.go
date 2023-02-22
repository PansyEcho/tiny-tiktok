package logic

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiny-tiktok/common/constant"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/common/utils"
	"tiny-tiktok/service/user/internal/model"
	"tiny-tiktok/service/user/internal/svc"
	"tiny-tiktok/service/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	resp = new(types.RegisterResp)
	user := new(model.User)

	result := l.svcCtx.DB.Where("username = ?", req.Username).Find(&user)
	if result.RowsAffected > 0 || result.Error != nil {
		registerCode(errx.DUPLICATE_USERNAME_ERROR, resp, err)
		return
	}

	user = &model.User{
		Username:        req.Username,
		Password:        utils.Md5(req.Password),
		Follow_count:    0,
		Follower_count:  0,
		Activity:        0,
		Avatar:          constant.DEFAULT_AVATAR_ADDRESS,
		BackgroundImage: constant.DEFAULT_BACKGROUND_ADDRESS,
		Signature:       "欢迎来到我的主页",
		TotalFavorited:  "0",
		WorkCount:       0,
		FavoriteCount:   0,
	}

	result = l.svcCtx.DB.Create(&user)
	if result.Error != nil {
		registerCode(errx.DB_ERROR, resp, err)
		return
	}

	generateToken, err := utils.GenerateToken(user.ID, user.Username, 60*60*24*7)
	if err != nil {
		registerCode(errx.TOKEN_GENERATE_ERROR, resp, err)
		return
	}

	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, generateToken, strconv.FormatInt(user.ID, 10), constant.PER_WEEK)
	if err != nil {
		registerCode(errx.REDIS_ERROR, resp, err)
		return
	}

	registerToken(errx.SUCCEED, resp, generateToken, user.ID)
	return
}

func registerCode(code int32, resp *types.RegisterResp, err error) {
	msg := errx.MapErrMsg(code)
	logx.Error(msg, err)
	err = errors.New(msg)
	resp.Status = types.Status{Status_msg: msg, Status_code: code}
}

func registerToken(code int32, resp *types.RegisterResp, token string, id int64) {
	msg := errx.MapErrMsg(code)
	resp.Status.Status_msg = msg
	resp.Status.Status_code = code
	resp.UserToken.Token = token
	resp.UserToken.UserID = id
}
