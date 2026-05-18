package model

import (
	"demo1/MyBlog/utils"
	"fmt"

	_ "github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
)

var db *gorm.DB
var err error
var (
	Redis *redis.Client
)

func InitDb() {
	db, err := gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName))
	if err != nil {
		fmt.Println("连接数据失败,请检查参数: ", err)
		return
	}

	db.SingularTable(true)
	db.AutoMigrate(&User{})
}

func InitRedis() {

}
