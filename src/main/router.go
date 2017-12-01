package main


import (
	"github.com/gin-gonic/gin"
	api "apis"
)


func initRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", api.IndexApi)

	router.POST("/person/create", api.AddPersonApi)

	router.GET("/person/get", api.GetPersonApi)
	//
	router.GET("/person/query", api.QueryPersonApi)
	//
	//router.PUT("/person/:id", ModPersonApi)
	//
	//router.DELETE("/person/:id", DelPersonApi)

	return router
}