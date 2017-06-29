package handlers

import (
	"github.com/Leafney/meiju-web/filters"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// router.RedirectTrailingSlash = true
	// router.RedirectFixedPath = true

	//全局过滤器
	router.Use(filters.Connect)
	router.Use(filters.ErrorHandler)

	//静态文件
	router.Static("/static", "./static")

	//Routers

	router.GET("/", Home_Index)

	//接口请求
	// api_Group := router.Group("/api")

	// api_v1 := api_Group.Group("/v1")

	// {
	// 	api_v1.GET("/videos", GetVideoList)

	// }

	return router
}
