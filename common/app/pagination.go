package app

import (
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return 10
	}
	if pageSize > 1000000 {
		return 1000000
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
