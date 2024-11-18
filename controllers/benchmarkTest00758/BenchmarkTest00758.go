package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00758Controller struct {
	web.Controller
}

func (c *BenchmarkTest00758Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00758Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00758", "")
	bar := "alsosafe"

	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[1]                                    // get the last 'safe' value
	}

	c.Ctx.Request.Context().Value("session").(http.ResponseWriter).Write([]byte(fmt.Sprintf("Item: 'userid' with value: '%s' saved in session.", bar)))

	output, err := json.Marshal(map[string]string{"userid": bar})
	if err == nil {
		c.Ctx.ResponseWriter.Write(output)
	}
}
