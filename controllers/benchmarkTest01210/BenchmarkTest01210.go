package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01210Controller struct {
	web.Controller
}

func (c *BenchmarkTest01210Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01210Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := c.Ctx.Request.Header["BenchmarkTest01210"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Error processing request.")
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Error processing request.")
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query("foo")
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Error processing request.")
		return
	}
	defer rows.Close()

	var results []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.CustomAbort(http.StatusInternalServerError, "Error processing request.")
			return
		}
		results = append(results, user)
	}

	output, err := json.Marshal(results)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Error processing request.")
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map86576 := make(map[string]interface{})
	map86576["keyA-86576"] = "a-Value"
	map86576["keyB-86576"] = param
	map86576["keyC"] = "another-Value"
	bar = map86576["keyB-86576"].(string)
	return bar
}
