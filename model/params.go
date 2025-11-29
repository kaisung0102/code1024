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
	Direction int   `json:"direction" binding:"required,oneof=1 0 -1"`
}
