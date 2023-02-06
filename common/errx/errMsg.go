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
}

func MapErrMsg(errcode int32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "请稍后再试... "
	}
}
