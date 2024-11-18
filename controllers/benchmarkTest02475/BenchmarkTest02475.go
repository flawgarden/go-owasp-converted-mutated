package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02475Controller struct {
	web.Controller
}

func (c *BenchmarkTest02475Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02475Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02475")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	hasher := md5.New()
	var input []byte
	input = []byte(bar)
	hasher.Write(input)

	result := hasher.Sum(nil)
	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

func doSomething(param string) string {
	bar := "safe!"
	map43776 := make(map[string]interface{})
	map43776["keyA-43776"] = "a-Value"
	map43776["keyB-43776"] = param
	map43776["keyC"] = "another-Value"
	bar = map43776["keyB-43776"].(string)

	return bar
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}
