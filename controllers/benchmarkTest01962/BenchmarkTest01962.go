package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01962Controller struct {
	web.Controller
}

func (c *BenchmarkTest01962Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01962Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	if benchmarkHeader := c.Ctx.Input.Header("BenchmarkTest01962"); benchmarkHeader != "" {
		param = benchmarkHeader
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	sqlStr := "SELECT * from USERS where USERNAME=? and PASSWORD='" + bar + "'"

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer statement.Close()

	_, err = statement.Exec("foo")
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.WriteString("Error processing request.")
		return
	}

	results, err := statement.Query()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer results.Close()

	users := []models.User{}
	for results.Next() {
		var user models.User
		if err := results.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			c.Ctx.WriteString("Error processing request.")
			return
		}
		users = append(users, user)
	}

	output, _ := json.Marshal(users)
	c.Ctx.WriteString(string(output))
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = string([]byte(param)) // This simulates some processing
	}

	return bar
}
