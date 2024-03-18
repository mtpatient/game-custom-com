package api

type FollowVo struct {
	Id        int    `json:"id"`
	Avatar    string `json:"avatar"`
	Username  string `json:"username"`
	Signature string `json:"signature"`
	IsFollow  int    `json:"is_follow"`
}

type FansVo struct {
	Id        int    `json:"id"`
	Avatar    string `json:"avatar"`
	Username  string `json:"username"`
	Signature string `json:"signature"`
	IsFollow  int    `json:"is_follow"`
}

type FansGet struct {
	Id        int `json:"id" v:"required"`
	PageIndex int `json:"page_index" v:"required"`
	PageSize  int `json:"page_size" v:"required"`
}
