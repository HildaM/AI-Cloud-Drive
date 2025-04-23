package user

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username" binding:"required,min=3,max=20"` // 用户名
	Password string `json:"password" binding:"required,min=8,max=30"` // 密码
}

// LoginRsp 登录响应
type LoginRsp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"` // 过期时间
	TokenType   string `json:"token_type"` // token类型
}
