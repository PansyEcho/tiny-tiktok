package errx

var message map[int32]string

func init() {
	message = make(map[int32]string)
	message[SUCCEED] = "SUCCESS"
	message[DUPLICATE_USERNAME_ERROR] = "用户名重复 "
	message[DB_ERROR] = "数据库发生错误 "
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆 "
	message[UNLOGIN_ERROR] = "请登陆后再试 "
	message[TOKEN_GENERATE_ERROR] = "生成token失败 "
	message[PASSWORD_ERROR] = "密码错误 "
	message[REDIS_ERROR] = "redis发生错误 "
	message[NOT_USER_ERROR] = "未知用户 "
	message[NOT_TOKEN_ERROR] = "未携带Token "
	message[UPDATE_ACTIVITY_ERROR] = "更新用户活跃度失败 "
	message[WRONG_DATA_ERROR] = "错误的视频数据 "
	message[UPLOAD_VIDEO_ERROR] = "上传视频失败 "
	message[UPLOAD_COVER_ERROR] = "上传封面失败 "
	message[NOT_PRIVILEGES_ERROR] = "无权限 "
	message[TRANSFORM_TIME_ERROR] = "时间转化错误 "
	message[REFOLLOW_ERROR] = "重复关注"
	message[KAFKAMARSHAL_ERROR] = "kafka序列化寄"
	message[KAFKASEND_ERROR] = "kafka发送信息寄"
	message[JSONMASHAL_ERROR] = "json序列化寄"
	message[KAFKAPUBLISH_ERROR] = "kafka发布寄"
	message[KAFKAPRODUCER_ERROR] = "kafka生产者寄"
	message[ACTIONTYPE_ERROR] = "ACTIONTYPE寄"
	message[UNFOLLOW_ERROR] = "未关注,无法取消"
	message[RPC_ERROR] = "RPC寄"
}

func MapErrMsg(errcode int32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "请稍后再试... "
	}
}
