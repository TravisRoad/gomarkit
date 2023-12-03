package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPageAndSize(c *gin.Context) (page int, size int) {
	p, ok := c.GetQuery("page")
	if ok {
		page, _ = strconv.Atoi(p)
	}
	s, ok := c.GetQuery("size")
	if ok {
		size, _ = strconv.Atoi(s)
	}
	if page < 0 {
		page = 1
	}
	if size < 10 {
		size = 10
	} else if size > 100 {
		size = 100
	}
	return
}
