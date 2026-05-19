package model

import (
	"demo1/MyBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
