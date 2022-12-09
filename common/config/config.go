package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"time"
)

/**
全局配置文件
*/

var (
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string

	GVA_DB *gorm.DB

	//AppMode  string
	HttpPort  string
	MediaUrl  string
	RootUrl   string
	RootMedia string

	JWT_SIGN_KEY    string
	JWT_EXPITE_TIME = time.Now().Add(7 * time.Hour).Unix()
)

// 初始化setting文件

func init() {
	file, err := ini.Load("common/config/config.ini") //你的ini文件所在的位置
	if err != nil {
		fmt.Println("配置文件错误无法读取", err)
	}
	LoadData(file)
}
func LoadData(file *ini.File) {
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	DB_HOST = file.Section("database").Key("DB_HOST").MustString("")
	DB_PORT = file.Section("database").Key("DB_PORT").MustInt(3306)
	DB_USER = file.Section("database").Key("DB_USER").MustString("")
	DB_PASSWORD = file.Section("database").Key("DB_PASSWORD").MustString("")
	DB_NAME = file.Section("database").Key("DB_NAME").MustString("")
	JWT_SIGN_KEY = file.Section("jwt").Key("SIGN_TOKEN").MustString("")
	MediaUrl = file.Section("server").Key("MEDIA_ROOT").MustString("")
	RootUrl = file.Section("server").Key("ROOT_URL").MustString("")
	RootMedia = file.Section("server").Key("ROOT_MEDIA").MustString("")
}
