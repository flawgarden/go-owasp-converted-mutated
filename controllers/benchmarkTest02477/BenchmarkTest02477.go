package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02477 struct {
	web.Controller
}

func (c *BenchmarkTest02477) Get() {
	c.post()
}

func (c *BenchmarkTest02477) Post() {
	c.post()
}

func (c *BenchmarkTest02477) post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02477")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", strings.ReplaceAll(bar, "'", "''"))
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	a40743 := param
	var b40743 strings.Builder
	b40743.WriteString(a40743)
	b40743.WriteString(" SafeStuff")
	b40743String := b40743.String()
	b40743String = strings.Replace(b40743String, b40743String[len(b40743String)-5:], "Chars", 1)
	c40743 := b40743String
	d40743 := c40743[:len(c40743)-1]

	return d40743
}
