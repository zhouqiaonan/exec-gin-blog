package routes

import (
	v1 "demo1/MyBlog/api/v1"
	"demo1/MyBlog/middleware"
	"demo1/MyBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
}
