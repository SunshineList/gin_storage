package main

import (
	common "gin_storage/common/config"
	globalinit "gin_storage/initialize"
)

func main() {

	g := globalinit.InitRoutes()             // 初始化路由
	common.GVA_DB = globalinit.DbInit()      // 初始化数据库
	globalinit.RegisterTables(common.GVA_DB) // 注册model
	globalinit.ZhValidatorInit()             // 注册验证翻译器

	g.Run(common.HttpPort)
}
