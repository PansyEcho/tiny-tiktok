package logic

import (
	"context"
	"fmt"
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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	resp = new(types.LoginResp)
	user := new(model.User)
	result := l.svcCtx.DB.Where("username = ? and password = ?", req.Username, utils.Md5(req.Password)).Limit(1).Find(&user)

	if result.RowsAffected == 0 {
		returnCode(errx.PASSWORD_ERROR, resp, err)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, constant.PER_WEEK)
	if err != nil {
		returnCode(errx.TOKEN_GENERATE_ERROR, resp, err)
		return
	}
	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, token, strconv.FormatInt(user.ID, 10), constant.PER_WEEK)
	if err != nil {
		returnCode(errx.REDIS_ERROR, resp, err)
		return
	}
	returnToken(errx.SUCCEED, resp, token, user.ID)
	fmt.Println("登录成功....")
	return
}

func returnCode(code int32, resp *types.LoginResp, err error) {
	msg := errx.MapErrMsg(code)
	logx.Error(msg, err)
	err = errors.New(msg)
	resp.Status = types.Status{Status_msg: msg, Status_code: code}
}

func returnToken(code int32, resp *types.LoginResp, token string, id int64) {
	msg := errx.MapErrMsg(code)
	resp.Status.Status_msg = msg
	resp.Status.Status_code = code
	resp.UserToken.Token = token
	resp.UserToken.UserID = id
}
