package model

import (
	"demo1/MyBlog/proto"
	"demo1/MyBlog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string    `gorm:"type:varchar(20);not null)" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email    string    `gorm:"type:varchar(32);not null" json:"email" validate:"required,email" label:"邮箱"`
	Role     int       `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"`
	Article  []Article `gorm:"many2many:user_article"`
	Status   string    `gorm:"type:varchar(12)" default:"N"`
	Code     string    `gorm:"type:varchar(80)"`
}

func CheckUser(name string) (code int) {
	var user User
	db.Model(&user).Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USERD
	}
	return errmsg.SUCCESS
}

func GetUserInfo(id int) (User, int) {
	var user User

	err = db.Model(&user).Preload("Article").Where("id = ?", id).First(&user).Error
	if user.ID > 0 {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

func EditUser(id int, data *proto.ReqEditUser) int {
	var maps = make(map[string]interface{})

	maps["username"] = data.Username
	maps["email"] = data.Email
	maps["role"] = data.Role
	err = db.Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
