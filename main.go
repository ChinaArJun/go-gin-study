package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github/go-gin-study/conf"
	"github/go-gin-study/corn"
	"net/http"
)

func main() {
	// 初始化配置
	conf.Init()

	// 初始化定时任务
	corn.StartCorn()

	fmt.Println("http_port :", viper.GetInt("HTTP_PORT"))
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	http.ListenAndServe(":8080", router)
	//s := &http.Server{
	//	Addr: ":8000",
	//	Handler: router,
	//	ReadHeaderTimeout: viper.GetDuration("READ_TIMEOUT"),
	//	WriteTimeout: viper.GetDuration("WRITE_TIMEOUT"),
	//	MaxHeaderBytes: 1 << 20,
	//}
	//// 运行服务
	//s.ListenAndServe()
}
