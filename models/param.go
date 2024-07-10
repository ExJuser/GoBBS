package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册的请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录的请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 谁向哪个帖子投了什么票
type ParamVoteData struct {
	//UserID 可以通过context获取
	PostID int64 `json:"post_id,string" binding:"required"`
	//赞成为1、反对为-1、0为取消
	Direction int8 `json:"direction,string" binding:"oneof=-1 0 1"`
}

type ParamPostList struct {
	Page  int64  `form:"page"`
	Size  int64  `form:"size"`
	Order string `form:"order"`
}
