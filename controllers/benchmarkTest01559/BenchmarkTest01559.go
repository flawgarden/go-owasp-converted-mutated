package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01559Controller struct {
	web.Controller
}

func (c *BenchmarkTest01559Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01559Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01559")
	if param == "" {
		param = ""
	}

	bar := new(Test).DoSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	results, err := statement.Exec()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	rowsAffected, _ := results.RowsAffected()
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Rows affected: %d", rowsAffected)))
}

type Test struct{}

func (t *Test) DoSomething(request *http.Request, param string) string {
	bar := "safe!"
	map31651 := make(map[string]interface{})
	map31651["keyA-31651"] = "a-Value"
	map31651["keyB-31651"] = param
	map31651["keyC"] = "another-Value"
	bar = map31651["keyB-31651"].(string)

	return bar
}
