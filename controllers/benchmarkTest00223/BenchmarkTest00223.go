package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type HashTestController struct {
	beego.Controller
}

func (c *HashTestController) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("some-header")
	if param == "" {
		c.Ctx.ResponseWriter.Write([]byte("Header not provided"))
		return
	}

	bar := "safe!"
	mapData := map[string]interface{}{
		"keyA": "a-Value",
		"keyB": param,
		"keyC": "another-Value",
	}
	bar = mapData["keyB"].(string)

	hashValue := hashFunction([]byte(bar))

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", bar)))
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Hash Test executed, value: %x\n", hashValue)))
}

func hashFunction(input []byte) []byte {
	// Simulate hashing for the purpose of the test
	return input // In real code, you would use a hashing algorithm here
}
