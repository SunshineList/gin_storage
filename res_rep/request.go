package request

/*
用户登录参数
*/

type LoginParams struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	CaptchaId  string `json:"captcha_id" binding:"required"`
	CaptchaVal string `json:"captcha_val" binding:"required"`
}

/*
	用户注册参数
*/

type RegisterParams struct {
	Username  string `json:"username" binding:"required"`
	Name      string `json:"name"`
	Password  string `json:"password" binding:"required"`
	Rpassword string `json:"rpassword" binding:"required"`
	Sex       *uint  `json:"sex"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
}

/*
	修改用户信息
*/

type UpdateUserInfoParams struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Sex      *uint  `json:"sex"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
}

/*
修改密码参数
*/
type updatePasswordParams struct {
	Password  string `json:"password" binding:"required"`
	Rpassword string `json:"rpassword" binding:"required"`
}
