package api

/*
	在这里定义好API的结构体 对外提供
*/

type RouterGroup struct {
	LoginApi
	UploadApi
}

var RouterGroupApp = new(RouterGroup)
