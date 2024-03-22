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
	TokenKey        = "game:token:"
	TokenKeyTTL     = 1800
	VerifyCodeKey   = "game:verifyCode:"
	PostLikesKey    = "game:like:post:"
	PostCollectKey  = "game:collect:post:"
	CommentLikesKey = "game:like:comment:"
	SearchKey       = "game:search:keywords"
)

const (
	LogSuccess = "success"
	LogError   = "error"
)
