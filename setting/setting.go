package setting

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	GinMode    string
	HttpPort   string
	Db         string
	DbName     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	AccessKey  string
	SecretKey  string
	Bucket     string
	CDNDomain  string
)

func Load() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Println("Fail to read file:", err)
		return
	}
	GinMode = cfg.Section("server").Key("gin_mode").String()
	HttpPort = cfg.Section("server").Key("http_port").String()
	Db = cfg.Section("database").Key("db").String()
	DbName = cfg.Section("database").Key("db_name").String()
	DbHost = cfg.Section("database").Key("db_host").String()
	DbPort = cfg.Section("database").Key("db_port").String()
	DbUser = cfg.Section("database").Key("db_user").String()
	DbPassword = cfg.Section("database").Key("db_password").String()
	AccessKey = cfg.Section("qiniu").Key("access_key").String()
	SecretKey = cfg.Section("qiniu").Key("secret_key").String()
	Bucket = cfg.Section("qiniu").Key("bucket").String()
	CDNDomain = cfg.Section("qiniu").Key("cdn_domain").String()
	fmt.Println("加载配置文件成功~~~")
}
