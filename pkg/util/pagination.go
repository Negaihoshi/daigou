package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"github.com/negaihoshi/daigou/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppConfig.App.PageSize
	}

	return result
}
