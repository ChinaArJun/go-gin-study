package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	//router := gin.Default()
	//router.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"messge": "ok",
	//	})
	//})
	//router.Run(":8081")
	
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	s := &http.Server{
		Addr: fmt.Sprintf(":%v", viper.GetString("HTTP_PORT")),
	}
	// 运行服务
	s.ListenAndServe()
}
