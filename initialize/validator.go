package initialize

import (
	"fmt"
	"gin_storage/utils"
)

func ZhValidatorInit() {
	if err := utils.LoadValidatorLocal("zh"); err != nil {
		fmt.Printf("验证翻译器初始化失败:%v\n", err)
		return
	}
}
