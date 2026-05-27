package v1

import (
	"demo1/MyBlog/model"
	"demo1/MyBlog/proto"
	"demo1/MyBlog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {
	var req proto.ReqEditUser
	var code int

	_ = c.ShouldBindJSON(&req)
	user, _ := model.GetUserInfo(req.Id)
	if user.Username != req.Username {
		code = model.CheckUser(req.Username)
		if code == errmsg.SUCCESS {
			c.JSON(code, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			return
		}
	}
	code = model.EditUser(req.Id, &req)

	profileOld, _ := model.GetProfileById(req.Id)
	profile := &model.Profile{
		ID:     req.Id,
		Name:   req.Username,
		Email:  req.Email,
		Desc:   profileOld.Desc,
		QqChat: profileOld.QqChat,
		WeChat: profileOld.WeChat,
		Weibo:  profileOld.Weibo,
		Img:    profileOld.Img,
		Avatar: profileOld.Avatar,
	}
	code = model.UpdateProfile(c, profile.ID, profile)
}
