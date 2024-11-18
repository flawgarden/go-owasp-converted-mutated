package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00769Controller struct {
	web.Controller
}

func (c *BenchmarkTest00769Controller) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest00769Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00769")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := "safe!"
	map29572 := map[string]interface{}{
		"keyA-29572": "a-Value",
		"keyB-29572": param,
		"keyC":       "another-Value",
	}
	bar = map29572["keyB-29572"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Abort("500")
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}

	c.Ctx.WriteString("No results can be displayed for query: " + sqlStr + "<br> because the Query method doesn't return results.")
}
