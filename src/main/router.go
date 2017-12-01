package main


import (
	"github.com/gin-gonic/gin"
	api "apis"
)


func initRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", api.IndexApi)
	router.POST("/user/create", api.AddUserApi)
	router.GET("/user/get", api.GetUserApi)
	router.GET("/user/query", api.QueryUserApi)
	router.POST("/user/update", api.ModUserApi)
	//router.DELETE("/person/delete", DelUserApi)

	return router
}