package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	ServerHost   string
	ServerPort   string
	FromEmail    string
	FromPassword string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查文件路径: ", err)
	}

	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
	LoadRedis(file)
	LoadEmailServer(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniuyun").Key("AccessKey").MustString("")
	SecretKey = file.Section("qiniuyun").Key("SecretKey").MustString("")
	Bucket = file.Section("qiniuyun").Key("Bucket").MustString("")
	QiniuServer = file.Section("qiniuyun").Key("Server").MustString("")
}

func LoadRedis(file *ini.File) {
	var redisSection = file.Section("redis")
	RedisAddr = redisSection.Key("RedisAddr").MustString("")
	RedisPassword = redisSection.Key("RedisPassword").MustString("")
	RedisDB = redisSection.Key("RedisDB").MustInt(0)
}

func LoadEmailServer(file *ini.File) {
	var emailSection = file.Section("email")
	ServerHost = emailSection.Key("ServerHost").MustString("")
	ServerPort = emailSection.Key("ServerPort").MustString("")
	FromEmail = emailSection.Key("FromEmail").MustString("")
	FromPassword = emailSection.Key("FromPassword").MustString("")
}
