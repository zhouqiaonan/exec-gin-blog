package routes

import (
	"demo1/MyBlog/api/v1"
	"demo1/MyBlog/middleware"
	"demo1/MyBlog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")
	r.StaticFile("admin/favicon.ico", "static/admin/favicon.ico")
	r.GET("admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.POST("user/update", v1.EditUser)
	}
}
