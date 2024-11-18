package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00931 struct {
	web.Controller
}

func (c *BenchmarkTest00931) Get() {
	c.Post()
}

func (c *BenchmarkTest00931) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00931")

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	responseMessage := fmt.Sprintf("No results can be displayed for query: %s<br>because the Spring execute method doesn't return results.", sqlStr)
	c.Ctx.ResponseWriter.Write([]byte(responseMessage))
}

func main() {
	web.Router("/sqli-01/BenchmarkTest00931", &BenchmarkTest00931{})
	web.Run()
}
