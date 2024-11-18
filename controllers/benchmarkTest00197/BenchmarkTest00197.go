package controllers

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00197Controller struct {
	web.Controller
}

func (c *BenchmarkTest00197Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00197Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest00197")
	param, _ = url.QueryUnescape(param)

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[1] // get the last 'safe' value
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	results := []string{}
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			http.Error(c.Ctx.ResponseWriter, "Error processing query.", http.StatusInternalServerError)
			return
		}
		results = append(results, username)
	}

	c.Ctx.ResponseWriter.Write([]byte("Your results are: <br>"))
	for _, s := range results {
		c.Ctx.ResponseWriter.Write([]byte(html.EscapeString(s) + "<br>"))
	}
}
