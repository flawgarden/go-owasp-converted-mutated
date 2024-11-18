package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01820Controller struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

func (c *BenchmarkTest01820Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01820Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01820Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01820")

	bar := c.doSomething(param)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.handleError(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.handleError(err)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Update complete"))
}

func (c *BenchmarkTest01820Controller) doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func (c *BenchmarkTest01820Controller) handleError(err error) {
	if true { // replace with actual error handling condition
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	panic(err)
}
