package controllers

import (
	"fmt"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00028Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00028Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00028Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00028")

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

list787231 := make([] string, 0)
list787231 = append(list787231, param)
list787231 = nil
param = list787231[0]

	if param != "" {
		fileName = "/path/to/testfiles/" + param

		var err error
		fos, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
			return
		}
		_, err = fos.WriteString("Now ready to write to file: " + htmlEscape(fileName))
		if err != nil {
			fmt.Printf("Error writing to file: %s\n", err)
		}
	}
}

func htmlEscape(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&", "&amp;"), "<", "&lt;")
}