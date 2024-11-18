package controllers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01243 struct {
	web.Controller
}

func init() {
	sql.Open("mysql", source)
}

func (c *BenchmarkTest01243) Get() {
	c.Post()
}

func (c *BenchmarkTest01243) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01243")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	filter := fmt.Sprintf("(&(objectclass=person)(uid=%s))", bar)

	results, err := db.Query("SELECT uid, street FROM users WHERE uid = ?", bar)
	if err != nil {
		panic(err)
	}
	defer results.Close()

	found := false
	for results.Next() {
		var uid, street string
		if err := results.Scan(&uid, &street); err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("LDAP query results:<br>Record found with name %s<br>Address: %s<br>", uid, street)))
		found = true
	}
	if !found {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("LDAP query results: nothing found for query: %s", htmlEscape(filter))))
	}
}

func htmlEscape(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "&", "&amp;"), "<", "&lt;")
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	if (7*42)-106 > 200 {
		bar = "This should never happen"
	}
	return bar
}
