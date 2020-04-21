package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github/go-gin-study/conf"
)

func GetPage(c *gin.Context) int { // 获取分页参数
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * conf.PageSize
	}

	return result
}