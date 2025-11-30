package model

type ParamsSignUp struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamsLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamsVote struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int    `json:"direction" binding:"oneof=1 0 -1"`
}

type ParamsGetPostList struct {
	Offset int64  `json:"offset" form:"offset"`
	Limit  int64  `json:"limit" form:"limit"`
	Order  string `json:"order" form:"order"`
}
