package utility

type R struct {
	Code int                    `json:"code" dc:"响应码"`
	Msg  string                 `json:"msg" dc:"响应信息"`
	Data map[string]interface{} `json:"data" dc:"响应数据"`
}

func GetR() *R {
	return &R{
		Code: 0,
		Msg:  "",
		Data: make(map[string]interface{}),
	}
}

func (r *R) Error(code int, msg string) *R {
	r.Code = code
	r.Msg = msg

	return r
}

func (r *R) PUT(key string, value interface{}) *R {
	r.Data[key] = value

	return r
}

func (r *R) SetCode(code int) *R {
	r.Code = code

	return r
}

func (r *R) SetMsg(msg string) *R {
	r.Msg = msg

	return r
}
