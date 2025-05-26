package dto

type UserRegistReq struct {
	UserName string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

type UserLoginReq struct {
	UserName string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

type UserLoginRes struct {
	Token string `json:"token"`
}

func NewUserLoginRes(token string) *UserLoginRes {
	return &UserLoginRes{
		Token: token,
	}
}
