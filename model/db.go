package model

import (
	"demo1/MyBlog/utils"
	"fmt"
	"time"

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

	//禁用默认表的复数形式
	db.SingularTable(true)
	//迁移: 创建表或者更新表
	db.AutoMigrate(&User{}, &Article{}, &Comment{}, &UserArticle{}, &Profile{})

	//设置连接池的最大闲置连接数
	db.DB().SetMaxIdleConns(10)
	//设置连接池的最大连接数量
	db.DB().SetMaxOpenConns(100)
	//设置连接的最大复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

}

func InitRedis() {

}
