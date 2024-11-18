package controllers

import (
	"fmt"
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00136Controller struct {
	web.Controller
}

func (c *BenchmarkTest00136Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00136Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest00136")
	param = strings.TrimSpace(param)

	bar := ""
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	fileName := fmt.Sprintf("testfiles/%s", bar)
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fos, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: ", fileName)
		return
	}

	_, err = c.Ctx.ResponseWriter.Write([]byte("Now ready to write to file: " + fileName))
	if err != nil {
		fmt.Println("Error writing response: ", err)
	}
}
