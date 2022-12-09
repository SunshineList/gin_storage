package middleware

import (
	"errors"
	"gin_storage/service"
	"github.com/gin-gonic/gin"
)

// 返回当前登录用户

func CurrentUser(context *gin.Context) (u interface{}, err error) {
	userId, exists := context.Get("user_id") // 这个在登陆签发jwt的时候会塞进上下文
	if exists == false {
		return nil, errors.New("未找到登录用户")
	}
	user, err := service.ServiceGroupApp.GetUserById(userId)
	return user, nil
}
