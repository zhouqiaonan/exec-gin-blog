package model

import ()

type UserArticle struct {
	ArticleId uint `gorm:"type:not null"  json:"article_id"`
	UserId    uint `gorm:"type:not null"  json:"user_id"`
}
