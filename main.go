package main

import (
	"demo1/MyBlog/model"
	"demo1/MyBlog/routes"
)

func main() {
	model.InitDb()
	model.InitRedis()
	routes.InitRouter()
}
