package model

import (
	"demo1/MyBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Commentator string `gorm:"type:varchar(255);not null" json:"commentator"`
	Content     string `gorm:"type:varchar(255);not null" json:"content"`
	ArticleID   int    `gorm:"type:int;not null" json:"article_id"`
	ParentID    int    `gorm:"type:int" json:"parent_id"`
}

func DeleteComment(id uint) int {
	err = db.Where("id = ?", id).Delete(&Comment{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
