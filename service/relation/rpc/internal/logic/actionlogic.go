package logic

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiny-tiktok/common/errx"
	"tiny-tiktok/service/relation/internal/types"
	"tiny-tiktok/service/relation/model/model"

	"tiny-tiktok/service/relation/rpc/internal/svc"
	"tiny-tiktok/service/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionLogic) Action(in *relation.ActionReq) (*relation.ActionResp, error) {
	// todo: add your logic here and delete this line
	switch in.Actiontype {
	case "1":
		flag, err := l.svcCtx.RedisCache.Sismember(string(in.Fromid), in.Followid)
		if err != nil {
			actionCode(errx.REDIS_ERROR, nil, err)
			return &relation.ActionResp{
				StatusMsg:  errx.MapErrMsg(errx.REDIS_ERROR),
				StatusCode: errx.REDIS_ERROR,
			}, err
		}
		if num, _ := l.svcCtx.RedisCache.Scard(string(in.Fromid)); num == 0 {
			var followlist []interface{}
			l.svcCtx.DB.Where("id = ?", in.Fromid).Select("follow_id").Find(&followlist)
			for _, k := range followlist {
				l.svcCtx.RedisCache.Sadd(string(in.Fromid), k)
			}
		}
		if flag {
			//重复关注错误
			actionCode(errx.REFOLLOW_ERROR, nil, err)
			return &relation.ActionResp{
				StatusMsg:  errx.MapErrMsg(errx.REFOLLOW_ERROR),
				StatusCode: errx.REFOLLOW_ERROR,
			}, err
		} else {
			//redis中操作
			l.svcCtx.RedisCache.Sadd(string(in.Fromid), in.Followid)
			fid, _ := strconv.Atoi(in.Followid)
			follow := &model.Follow{
				UserId:   in.Fromid,
				FollowId: int64(fid),
			}
			result := l.svcCtx.DB.Create(&follow)
			if result.Error != nil {
				actionCode(errx.DB_ERROR, nil, err)
				return &relation.ActionResp{
					StatusMsg:  errx.MapErrMsg(errx.DB_ERROR),
					StatusCode: errx.DB_ERROR,
				}, err
			} else {
				return &relation.ActionResp{
					StatusMsg:  "关注成功",
					StatusCode: 0,
				}, nil
			}
			//go-queue版本还没整好
			/*
				key := fmt.Sprintf("follow:%d:%d", in.Fromid, in.Followid)
				err := l.svcCtx.RedisCache.Set(key, "1")
				l.svcCtx.RedisCache.Expire(key, 1000*30)
				if err != nil {
					actionCode(errx.REDIS_ERROR, nil, err)
					return nil, err
				} else {
					producer, err := kqueue.GetProducer()
					if err != nil {
						actionCode(errx.KAFKAPRODUCER_ERROR, resp, err)
						return nil, err
					}
					err = relation_producer.Relation(producer, relation_producer.RelationMessage{
						FromId:     string(userClaim.Id),
						FollowId:   req.To_user_id,
						ActionType: req.Action_type,
					})
					if err != nil {
						actionCode(errx.KAFKAMARSHAL_ERROR, resp, err)
						return
					}
					return &types.ActionResp{
						types.Status{
							Status_code: 0,
							Status_msg:  "relation successfully",
						},
					}, nil
				}*/
		}
		break
	case "2":
		flag, err := l.svcCtx.RedisCache.Sismember(string(in.Fromid), in.Followid)
		if err != nil {
			actionCode(errx.REDIS_ERROR, nil, err)
			return &relation.ActionResp{
				StatusMsg:  errx.MapErrMsg(errx.REDIS_ERROR),
				StatusCode: errx.REDIS_ERROR,
			}, err
		}
		if num, _ := l.svcCtx.RedisCache.Scard(string(in.Fromid)); num == 0 {
			var followlist []interface{}
			l.svcCtx.DB.Where("id = ?", in.Fromid).Select("follow_id").Find(&followlist)
			for _, k := range followlist {
				l.svcCtx.RedisCache.Sadd(string(in.Fromid), k)
			}
		}
		if flag {
			l.svcCtx.RedisCache.Srem(string(in.Fromid), in.Followid)
			fid, _ := strconv.Atoi(in.Followid)
			follow := &model.Follow{
				UserId:   in.Fromid,
				FollowId: int64(fid),
			}
			result := l.svcCtx.DB.Delete(&follow)
			if result.Error != nil {
				actionCode(errx.DB_ERROR, nil, err)
				return &relation.ActionResp{
					StatusMsg:  errx.MapErrMsg(errx.DB_ERROR),
					StatusCode: errx.DB_ERROR,
				}, err
			} else {
				return &relation.ActionResp{
					StatusMsg:  "取关成功",
					StatusCode: 0,
				}, nil
			}
			//go-queue版本还没整好
			/*
				key := fmt.Sprintf("follow:%d:%d", userClaim.Id, req.To_user_id)
				err := l.svcCtx.RedisCache.Set(key, "1")
				l.svcCtx.RedisCache.Expire(key, 1000*30)
				if err != nil {
					actionCode(errx.REDIS_ERROR, resp, err)
					return
				} else {
					producer, err := kqueue.GetProducer()
					if err != nil {
						actionCode(errx.KAFKAPRODUCER_ERROR, resp, err)
						return
					}
					err = relation_producer.Relation(producer, relation_producer.RelationMessage{
						FromId:     string(userClaim.Id),
						FollowId:   req.To_user_id,
						ActionType: req.Action_type,
					})
					if err != nil {
						actionCode(errx.KAFKAMARSHAL_ERROR, resp, err)
						return
					}
					return &types.ActionResp{
						types.Status{
							Status_code: 0,
							Status_msg:  "relation successfully",
						},
					}, nil
				}*/
		} else {
			actionCode(errx.UNFOLLOW_ERROR, nil, err)
			return &relation.ActionResp{
				StatusMsg:  errx.MapErrMsg(errx.UNFOLLOW_ERROR),
				StatusCode: errx.UNFOLLOW_ERROR,
			}, err
		}
		break
	default:
		actionCode(errx.ACTIONTYPE_ERROR, nil, errors.New(errx.MapErrMsg(errx.ACTIONTYPE_ERROR)))
		return &relation.ActionResp{
			StatusMsg:  errx.MapErrMsg(errx.ACTIONTYPE_ERROR),
			StatusCode: errx.ACTIONTYPE_ERROR,
		}, errors.New(errx.MapErrMsg(errx.ACTIONTYPE_ERROR))
	}
	return &relation.ActionResp{}, nil
}
func actionCode(code int32, resp *types.ActionResp, err error) {
	msg := errx.MapErrMsg(code)
	logx.Error(msg, err)
	err = errors.New(msg)
	resp.Status = types.Status{Status_msg: msg, Status_code: code}
}
