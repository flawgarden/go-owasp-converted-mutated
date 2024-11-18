package controllers

import (
	"fmt"
	"net/url"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00132 struct {
	web.Controller
}

func (c *BenchmarkTest00132) Get() {
	c.Post()
}

func (c *BenchmarkTest00132) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Input.Header("BenchmarkTest00132")

	// URL Decode the header value
	decodedParam, _ := url.QueryUnescape(param)

	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B'

	switch switchTarget {
	case 'A':
		bar = decodedParam
	case 'B':
		bar = "bob"
	case 'C':
	case 'D':
		bar = decodedParam
	default:
		bar = "bob's your uncle"
	}

	fileName := fmt.Sprintf("/path/to/test/files/%s", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Problem getting FileInputStream: %s", err.Error())))
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("The beginning of file: '%s' is:\n\n%s", fileName, string(b[:size]))))
}

func main() {
	web.Router("/pathtraver-00/BenchmarkTest00132", &BenchmarkTest00132{})
	web.Run()
}
