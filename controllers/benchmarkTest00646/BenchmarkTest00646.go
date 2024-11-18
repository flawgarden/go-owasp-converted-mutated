package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00646Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00646Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00646Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00646Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00646")
	if param == "" {
		param = ""
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[1]                                    // get the last 'safe' value
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output, err := json.Marshal(bar)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}
