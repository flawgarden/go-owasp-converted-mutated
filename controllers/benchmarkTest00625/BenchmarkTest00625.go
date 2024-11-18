package controllers

import (
	"fmt"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00625Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00625Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00625Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00625")
	if param == "" {
		param = ""
	}

	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = fmt.Sprintf("./%s", bar) // Update path as needed
	var err error
	fos, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", fileName)))
}
