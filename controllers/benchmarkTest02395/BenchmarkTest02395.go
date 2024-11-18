package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02395Controller struct {
	web.Controller
}

func (c *BenchmarkTest02395Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02395Controller) Post() {
	id := c.GetString("BenchmarkTest02395")
	if id == "" {
		id = ""
	}

	bar := doSomething(id)

	response := map[string]string{"result": fmt.Sprintf("Formatted like: %s.", bar)}
	output, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func main() {
	web.Router("/xss-04/BenchmarkTest02395", &BenchmarkTest02395Controller{})
	web.Run()
}
