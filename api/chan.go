package api

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	for i := 0; i < 1000; i++ {
		ch := make(chan int, 50)

		ch = ch
	}

	c.String(200, "ok")
}