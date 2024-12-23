package controllers

import (
"encoding/base64"
"fmt"
"html"
"os"
"github.com/beego/beego/v2/server/web"
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

generator := makeMessageGeneratorBroken(param)()
param = generator("fixed_string")

	if param != "" {
		decoded, err := base64.StdEncoding.DecodeString(param)
		if err == nil {
			bar = string(decoded)
		}
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

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


