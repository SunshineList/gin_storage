package utils

import (
	"github.com/mojocn/base64Captcha"
)

/**
验证器相关代码
*/

// 验证码校验

func CheckCaptcha(captchaId string, captchaVal string) bool {
	var store = base64Captcha.DefaultMemStore
	return store.Verify(captchaId, captchaVal, true)
}
