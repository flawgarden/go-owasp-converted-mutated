package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02628Controller struct {
	web.Controller
}

func (c *BenchmarkTest02628Controller) Get() {
	c.post(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest02628Controller) Post() {
	c.post(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest02628Controller) post(req *http.Request, resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := req.URL.RawQuery
	paramval := "BenchmarkTest02628="
	paramLoc := -1
	if queryString != "" {
		paramLoc = strings.Index(queryString, paramval)
	}
	if paramLoc == -1 {
		resp.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02628")))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, err := url.QueryUnescape(param)
	if err != nil {
		http.Error(resp, "Error decoding parameter", http.StatusBadRequest)
		return
	}

	bar := doSomething(param)
	sqlStr := fmt.Sprintf("CALL %s", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(resp, "Error opening database connection", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(resp, "Error executing query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(resp, "Error scanning results", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(resp, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}
	resp.Write(output)
}

func doSomething(param string) string {
	bar := "safe!"
	map23653 := map[string]interface{}{
		"keyA-23653": "a-Value",
		"keyB-23653": param,
		"keyC":       "another-Value",
	}
	bar = map23653["keyB-23653"].(string)
	return bar
}
