package api

import (
	"gin_storage/common/config"
	"gin_storage/middleware"
	"gin_storage/model"
	"gin_storage/res_rep"
	"gin_storage/response"
	"gin_storage/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strconv"
)

type LoginApi struct{}

func (l *LoginApi) Login(context *gin.Context) {

	var loginParams = request.LoginParams{}
	err := context.ShouldBindJSON(&loginParams)
	if err != nil {
		response.FailAndMsg(utils.Translate(err), context)
		return
	}
	user := &model.User{
		Username: loginParams.Username,
		Password: loginParams.Password,
	}

	if !utils.CheckCaptcha(loginParams.CaptchaId, loginParams.CaptchaVal) {
		response.FailAndMsg("验证码错误", context)
		return
	}

	res, err := userService.LoginService(*user)

	if err != nil {
		response.FailAndMsg(err.Error(), context)
		return
	}

	if res.Status != model.ACTIVE {
		response.FailAndMsg("账号已冻结请联系管理员", context)
		return
	}

	// 签发jwt
	jwtToken, _ := utils.JwtToken(utils.MyCustomClaims{Id: res.ID, Username: res.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: config.JWT_EXPITE_TIME, // 过期时间
		}})

	response.OkAndData(map[string]interface{}{
		"username": res.Username,
		"name":     res.Name,
		"token":    jwtToken,
	}, "登录成功", context)
}

func (l *LoginApi) Register(context *gin.Context) {
	var registerParams = request.RegisterParams{}
	err := context.ShouldBindJSON(&registerParams)
	if err != nil {
		response.FailAndMsg(utils.Translate(err), context)
		return
	}

	if registerParams.Password != registerParams.Rpassword {
		response.FailAndMsg("两次密码输入不一致", context)
		return
	}

	user := &model.User{
		Username: registerParams.Username,
		Name:     registerParams.Name,
		Password: registerParams.Password,
		Sex:      registerParams.Sex,
		Phone:    registerParams.Phone,
	}

	r, err := userService.RegisterService(*user)

	if err != nil {
		response.FailAndMsg("注册失败", context)
		return
	}
	response.OkAndData(r, "注册成功", context)
}

func (l *LoginApi) GetUserInfo(context *gin.Context) {
	user, err := middleware.CurrentUser(context)
	if err != nil {
		response.FailAndMsg(err.Error(), context)
		return
	}
	response.OkAndData(user, "查询成功", context)
}

func (l *LoginApi) UpdateUser(context *gin.Context) {
	var updateParams = request.UpdateUserInfoParams{}
	var userId = context.Params.ByName("id")
	err := context.ShouldBindJSON(&updateParams)
	if err != nil {
		response.FailAndMsg(utils.Translate(err), context)
		return
	}

	user_id, err := strconv.Atoi(userId)

	user := &model.User{
		BaseModel: model.BaseModel{
			ID: uint(user_id),
		},
		Username: updateParams.Username,
		Name:     updateParams.Name,
		Avatar:   updateParams.Avatar,
		Phone:    updateParams.Phone,
		Sex:      updateParams.Sex,
	}

	res, err := userService.UpdateUserService(*user)

	if err != nil {
		response.FailAndMsg("更新信息失败", context)
		return
	}
	response.OkAndData(res, "更新成功", context)
}
