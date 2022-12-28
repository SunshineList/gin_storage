package middleware

import (
	"errors"
	"fmt"
	"gin_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		fmt.Println(method)
		fmt.Println(origin)
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
