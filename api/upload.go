package api

import (
	"gin_storage/common/config"
	"gin_storage/response"
	"github.com/gin-gonic/gin"
)

type UploadApi struct{}

func (u *UploadApi) UploadFile(context *gin.Context) {
	urls := []string{}
	form, _ := context.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		//上传文件到指定的路径
		context.SaveUploadedFile(file, config.MediaUrl+file.Filename)
		urls = append(urls, config.RootMedia+file.Filename)
	}
	response.OkAndData(urls, "上传成功", context)
}
