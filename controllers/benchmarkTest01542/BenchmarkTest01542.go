package controllers

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Register driver and database
}

type BenchmarkTest01542Controller struct {
	web.Controller
}

func (c *BenchmarkTest01542Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01542Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01542Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01542")
	if param == "" {
		param = ""
	}


	// Here should be the logic handling SecureRandom that checks user presence
}

func (c *BenchmarkTest01542Controller) doSomething(param string) string {
	bar := "safe!"
	mapData := make(map[string]interface{})
	mapData["keyA-98601"] = "a-Value"
	mapData["keyB-98601"] = param
	mapData["keyC"] = "another-Value"
	bar = mapData["keyB-98601"].(string)

	return bar
}
