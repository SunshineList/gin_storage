package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 1
	ERROR   = 0
)

func Result(code int, data interface{}, msg string, content *gin.Context) {
	content.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	if code == ERROR {
		return
	}
}

//成功请求返回结构(带数据的)

func OkAndData(data interface{}, msg string, content *gin.Context) {
	Result(SUCCESS, data, msg, content)
}

// 只返回成功信息

func Ok(content *gin.Context) {
	//Result(SUCCESS, map[string]interface{}{}, content)
	Result(SUCCESS, nil, "操作成功", content)
}

// 失败返回结构体

func FailAndMsg(msg string, content *gin.Context) {
	//Result(ERROR, map[string]interface{}{}, msg, content)
	Result(ERROR, nil, msg, content)
}
