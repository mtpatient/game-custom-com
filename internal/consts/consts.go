package consts

const ContextKey = "user"

/*
*
错误码
*/
const (
	RequestErrCode = 450 // 请求解析错误
	ServiceErrCode = 550 // 业务层错误
	RedisErrCode   = 600
)

/*
*
redis 前缀
*/
const (
	TokenKey      = "game:token:"
	TokenKeyTTL   = 1800
	LikeKey       = "like:post:"
	UserKey       = "game:user:"
	VerifyCodeKey = "game:verifyCode:"
)

const (
	LogSuccess = "success"
	LogError   = "error"
)
