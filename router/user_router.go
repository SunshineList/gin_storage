package api

import (
	"gin_storage/api"
	"gin_storage/middleware"
	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

func (l *LoginApi) LoginRouters(Router *gin.RouterGroup) {
	r := Router.Group("user").Use(middleware.Cors()) //.Use(middleware.TokenMiddleware())
	{
		r.POST("/login", api.ApiGroupApp.Login)                                                  // 登录接口
		r.POST("/register", api.ApiGroupApp.Register)                                            //注册
		r.GET("/getCaptcha", api.ApiGroupApp.GetCaptcha)                                         // 获取验证码
		r.GET("/getUserInfo", middleware.JWTAuthMiddleware(), api.ApiGroupApp.GetUserInfo)       // 查看当前用户信息
		r.PUT("/updateUserInfo/:id", middleware.JWTAuthMiddleware(), api.ApiGroupApp.UpdateUser) // 修改用户信息
	}
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywidXNlcm5hbWUiOiJhZG1pbiIsImV4cCI6MTY2OTczNjgyMn0.UGKf0FrPb-dOGNHzo8KdbHLSneC3Ux44Tg_VquOHcPo
