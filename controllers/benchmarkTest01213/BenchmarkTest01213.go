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

func init() {
	// Database initialization code here
}

type BenchmarkTest01213Controller struct {
	web.Controller
}

func (c *BenchmarkTest01213Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01213Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	if header := c.Ctx.Input.Header("BenchmarkTest01213"); header != "" {
		param = header
	}

	param, _ = url.QueryUnescape(param)

	bar := c.doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Query("foo")
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer result.Close()

	user := models.User{}
	for result.Next() {
		err = result.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
			return
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *BenchmarkTest01213Controller) doSomething(param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
