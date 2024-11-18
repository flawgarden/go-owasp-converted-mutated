package controllers

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00626Controller struct {
	web.Controller
}

func (c *BenchmarkTest00626Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00626Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00626")
	if param == "" {
		param = ""
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = fmt.Sprintf("path/to/directory/%s", bar)

	fos, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", fileName)))
}
