package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01979Controller struct {
	web.Controller
}

func (c *BenchmarkTest01979Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01979Controller) Post() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	for name, values := range c.Ctx.Request.Header {
		if contains(commonHeaders, name) {
			continue
		}
		if len(values) > 0 {
			param = name
			break
		}
	}

	bar := doSomething(param)
	result, err := encrypt(bar)
	if err != nil {
		http.Error(response, "Encryption error", http.StatusInternalServerError)
		return
	}

	_, err = response.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", result)))
	if err != nil {
		http.Error(response, "Write error", http.StatusInternalServerError)
		return
	}
}

func doSomething(param string) string {
	return htmlEscape(param)
}

func encrypt(input string) (string, error) {
	// Implement encryption logic here
	return input, nil // Placeholder
}

var commonHeaders = []string{"Host", "User-Agent", "Accept", "Accept-Encoding", "Accept-Language"}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func htmlEscape(s string) string {
	return s // Implement HTML escaping logic here
}
