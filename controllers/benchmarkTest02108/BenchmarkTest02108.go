package controllers

import (
	"fmt"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02108Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02108Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02108Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02108")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = fmt.Sprintf("testfiles/%s", bar)

	var err error
	fos, err = os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", fileName)))
}

func doSomething(param string) string {
	bar := "safe!"
	map44255 := make(map[string]interface{})
	map44255["keyA-44255"] = "a_Value"
	map44255["keyB-44255"] = param
	map44255["keyC"] = "another_Value"
	bar = fmt.Sprintf("%s", map44255["keyB-44255"])
	bar = fmt.Sprintf("%s", map44255["keyA-44255"])

	return bar
}
