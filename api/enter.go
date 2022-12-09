package api

import "gin_storage/service"

type ApiGroup struct {
	LoginApi
	UploadApi
	CaptchaApi
}

var ApiGroupApp = new(ApiGroup)

var (
	userService = service.ServiceGroupApp.UserService
)
