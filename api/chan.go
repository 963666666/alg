package api

import (
	_ "embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	hello := []*user{}
	hello = append(hello, &user{
		Prefer: "hello",
	})

	helloBytes, _ := json.Marshal(hello)

	c.String(200, string(helloBytes))
}

type user struct {
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