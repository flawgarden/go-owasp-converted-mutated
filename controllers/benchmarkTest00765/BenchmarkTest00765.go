package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest00765Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00765Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00765Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00765Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00765", "")
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("No results returned for query: " + sqlStr))
		return
	}

	output, err := json.Marshal(map[string]interface{}{"results": results})
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	// Placeholder for real functionality
	return param
}
