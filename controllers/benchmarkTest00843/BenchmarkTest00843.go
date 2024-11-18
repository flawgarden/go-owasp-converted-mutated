package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln2Controller struct {
	web.Controller
}

func (c *SqlInjectionVuln2Controller) Get() {
	queryString := c.Ctx.Request.URL.Query()
	param := queryString.Get("BenchmarkTest00843")
	bar := "safe!"
	if param != "" {
		bar = param
	}

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.WriteString("Error connecting to database")
		return
	}
	defer db.Close()

	var userid int64
	err = db.QueryRow(sqlStr).Scan(&userid)
	if err != nil {
		c.Ctx.WriteString("No results returned for query")
		return
	}

	output, err := json.Marshal(userid)
	if err != nil {
		c.Ctx.WriteString("Error marshalling results")
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
