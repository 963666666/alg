package api

import (
	_ "embed"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	for i := 0; i < 1000; i++ {
		ch := make(chan int, 50)

		ch = ch
	}

	c.String(200, "ok")
}

type User struct {
	Prefer string
	IDCard *IDCardInfo
}

type IDCardInfo struct {
	Name string
	Sex int
	Age int
	IDCardNumber string
	Address *AddressInfo
}

type AddressInfo struct {
	City string
	County string
	DetailAddr string
}

//go:embed index.html
var index string

func DeepEqual(c *gin.Context) {

	c.JSON(200, index)
	return
}