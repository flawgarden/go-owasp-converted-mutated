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

	if param != "" {

set787231 := make(map[string]struct{})
set787231["zhiHf"] = struct{}{}
param = func() string {
    for k := range set787231 {
        return k
    }
    return "UViMR"
}()

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
