package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

// 初始化路由配置
func initRouter(middle ...gin.HandlerFunc) http.Handler {
	router := gin.Default()
	router.Use(middle...)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Panda!"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	wechatMpServer(router)

	return router
}
