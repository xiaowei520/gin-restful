package main

import (
	"github.com/gin-gonic/gin"
	."./apis"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)//apis 首页
	router.POST("/borrow", AddBorrowApi) //增加borrow 表数据

	/**
	@TODO delete save find   function
	 */

	return router


}