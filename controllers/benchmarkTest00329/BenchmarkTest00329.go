package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00329Controller struct {
	web.Controller
}

func (c *BenchmarkTest00329Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00329Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	if headers := c.Ctx.Request.Header["BenchmarkTest00329"]; len(headers) > 0 {
		param = headers[0]
	}
	param, _ = url.QueryUnescape(param)

	bar := param
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	}

	sqlStr := fmt.Sprintf("{call %s}", bar)

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

	rs, err := stmt.Query()
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rs.Close()

	var results []models.User
	for rs.Next() {
		var user models.User
		if err = rs.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
			return
		}
		results = append(results, user)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
