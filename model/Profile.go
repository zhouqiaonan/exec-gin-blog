package model

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
