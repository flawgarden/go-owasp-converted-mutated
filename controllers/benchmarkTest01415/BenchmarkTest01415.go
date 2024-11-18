package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01415 struct {
	web.Controller
}

func (c *BenchmarkTest01415) Get() {
	c.Post()
}

func (c *BenchmarkTest01415) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	names := c.Ctx.Request.URL.Query()
	for name := range names {
		values := c.Ctx.Request.URL.Query()[name]
		for _, value := range values {
			if value == "BenchmarkTest01415" {
				param = name
				break
			}
		}
		if param != "" {
			break
		}
	}

	bar := new(Test).doSomething(param)

	hashValue, err := hashAndStore(bar)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", bar)))
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Hash Test executed: %s", hashValue)))
}

func hashAndStore(input string) (string, error) {
	// Hashing logic similar to the one in Java goes here
	// For simplification, returning a mock hash value
	hashValue := "mocked_hash_value"
	return hashValue, nil
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param
	}
	return bar
}
