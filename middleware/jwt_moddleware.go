package middleware

import (
	"fmt"
	"gin_storage/response"
	"gin_storage/utils"
	"github.com/gin-gonic/gin"
)

// 基于JWT的认证中间件

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从请求头中取出
		signToken := context.Request.Header.Get("Authorization")
		if signToken == "" {
			response.FailAndMsg("token为空", context)
			context.Abort()
			return
		}
		// 校验token
		myclaims, err := utils.ParserToken(signToken)
		if err != nil {
			fmt.Println(err)
			response.FailAndMsg("token校验失败", context)
			context.Abort()
			return
		}
		// 将用户的id放在到请求的上下文中
		context.Set("user_id", myclaims.Id)
		context.Next() // 后续的处理函数可以用过context.Get("userid")来获取当前请求的id

	}

}
