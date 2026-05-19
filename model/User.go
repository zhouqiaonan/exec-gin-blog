package model

import (
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
