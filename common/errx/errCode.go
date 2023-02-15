package errx

const (
	SUCCEED int32 = 0
	ERROR   int32 = 1
)

// 全局错误码
const (
	DUPLICATE_USERNAME_ERROR int32 = iota + 7001
	DB_ERROR
	TOKEN_EXPIRE_ERROR
	TOKEN_GENERATE_ERROR
	PASSWORD_ERROR
	REDIS_ERROR
	UNLOGIN_ERROR
	NOT_USER_ERROR
	NOT_TOKEN_ERROR
	UPDATE_ACTIVITY_ERROR
	WRONG_DATA_ERROR
)
