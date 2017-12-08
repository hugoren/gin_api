package main


import (
	"github.com/gin-gonic/gin"
	api "apis"
)


func initRouter() *gin.Engine {

	router := gin.Default()

	// 首页
	router.GET("/", api.IndexApi)

	// curd示例
	router.POST("/user/create", api.AddUserApi)
	router.GET("/user/get", api.GetUserApi)
	router.GET("/user/query", api.QueryUserApi)
	router.POST("/user/update", api.ModUserApi)
	router.GET("/user/del", api.DelUserApi)

	// 索引查询
	router.POST("/index/query", api.IndexQuery)

	return router
}