package model

import (
	"demo1/MyBlog/utils"
	"demo1/MyBlog/utils/errmsg"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Profile struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Name   string `gorm:"type:varchar(20)" json:"name"`
	Desc   string `gorm:"type:varchar(200)" json:"desc"`
	QqChat string `gorm:"type:varchar(200)" json:"qq_chat"`
	WeChat string `gorm:"type:varchar(200)" json:"wechat"`
	Weibo  string `gorm:"type:varchar(200)" json:"weibo"`
	Email  string `gorm:"type:varchar(200)" json:"email"`
	Img    string `gorm:"type:varchar(200)" json:"img"`
	Avatar string `gorm:"type:varchar(200)" json:"avatar"`
}

func GetProfileById(id int) (*Profile, int) {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return &profile, errmsg.SUCCESS
}

func UpdateProfile(c *gin.Context, id int, profile *Profile) int {
	var _profile *Profile
	err := db.Model(&Profile{}).Where("ID = ?", id).Updates(profile).Error
	if err != nil {
		return errmsg.ERROR
	}
	sessionId, err := c.Cookie(SessionName)
	if err != nil {
		return errmsg.ERROR
	}
	result, err := Redis.Get(sessionId).Result()
	if err != nil {
		return errmsg.ERROR
	}
	err = json.Unmarshal([]byte(result), &_profile)
	if err != nil {
		fmt.Println("Unmarshal json false")
	}
	if _profile.ID == profile.ID {
		profileJson, _ := json.Marshal(profile)
		_, err = Redis.Set(sessionId, string(profileJson), utils.Expiration).Result()
		if err != nil {
			return errmsg.ERROR
		}
	}
	return errmsg.SUCCESS
}
