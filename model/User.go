package model

import (
	"demo1/MyBlog/proto"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null)" json:"password" validate:"required,min=6,max=20" label:"密码"`
}
