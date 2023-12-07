package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func clamp(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func getPageAndSize(c *gin.Context) (page int, size int) {
	p, ok := c.GetQuery("page")
	if ok {
		page, _ = strconv.Atoi(p)
	}
	s, ok := c.GetQuery("size")
	if ok {
		size, _ = strconv.Atoi(s)
	}
	if page <= 0 {
		page = 1
	}
	size = clamp(size, 10, 100)
	return
}
