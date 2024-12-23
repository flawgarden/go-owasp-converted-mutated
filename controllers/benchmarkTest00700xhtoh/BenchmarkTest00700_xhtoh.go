package controllers

import (
"encoding/base64"
"fmt"
"html"
"os"
"github.com/beego/beego/v2/server/web"
"sync"
)

type BenchmarkTest00700Controller struct {
	web.Controller
}

func (c *BenchmarkTest00700Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00700Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := c.GetString("BenchmarkTest00700")
	bar := ""

	if param != "" {
		decoded, err := base64.StdEncoding.DecodeString(param)
		if err == nil {
			bar = string(decoded)
		}

dataChannel := make(chan Data, 1)
dataChannel <- Data{Value: bar}

var wg sync.WaitGroup
wg.Add(1)

go func() {
    data := <-dataChannel
    data.Value += "suffix"
    dataChannel <- data
}()

wg.Wait()

readData := <-dataChannel
bar = readData.Value

	}

	fileName := fmt.Sprintf("testfiles/%s", bar)
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	if bar != "" {
		var err error
		fos, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
			return
		}
		c.Ctx.WriteString(fmt.Sprintf("Now ready to write to file: %s", html.EscapeString(fileName)))
	}
}
