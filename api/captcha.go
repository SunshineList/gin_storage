package api

import (
	"gin_storage/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaApi struct{}

var store = base64Captcha.DefaultMemStore

func (cp *CaptchaApi) GetCaptcha(context *gin.Context) {
	// height 高度 png 像素高度
	// width  宽度 png 像素高度
	// length 验证码默认位数
	// maxSkew 单个数字的最大绝对倾斜因子
	// dotCount 背景圆圈的数量
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, content, err := captcha.Generate()
	if err != nil {
		response.FailAndMsg("验证码生成失败", context)
		return
	}
	response.OkAndData(map[string]interface{}{
		"captchaId": id,
		"content":   content,
	}, "验证码生成成功", context)
}
