package api

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	a := 10

	defer func(i *int) {
		fmt.Println("hello defer : ", *i)
	}(&a)

	a += 50

	c.String(200, string("1"))
}

type user struct {
	Prefer string
	IDCard *IDCardInfo
}

type IDCardInfo struct {
	Name         string
	Sex          int
	Age          int
	IDCardNumber string
	Address      *AddressInfo
}

type AddressInfo struct {
	City       string
	County     string
	DetailAddr string
}

//go:embed index.html
var index string

func DeepEqual(c *gin.Context) {

	c.JSON(200, index)
	return
}
