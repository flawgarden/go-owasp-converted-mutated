package controllers

import (
	"crypto/sha512"
	"database/sql"
	"fmt"
	"html"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

type BenchmarkTest01037 struct {
	web.Controller
}

func (c *BenchmarkTest01037) Get() {
	c.doPost()
}

func (c *BenchmarkTest01037) Post() {
	c.doPost()
}

func (c *BenchmarkTest01037) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest01037")
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	input := []byte(bar)
	md := sha512.New()
	md.Write(input)

	// Simulate file writing (update as necessary)
	// fileTarget := "path/to/passwordFile.txt"
	// ioutil.WriteFile(fileTarget, []byte(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(result))), os.O_APPEND|os.O_CREATE|os.O_WRONLY)

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", html.EscapeString(string(input)))))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}
