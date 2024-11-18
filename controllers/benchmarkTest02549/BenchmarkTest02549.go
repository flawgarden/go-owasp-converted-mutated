package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02549Controller struct {
	web.Controller
}

func init() {
	web.Router("/crypto-02/BenchmarkTest02549", &BenchmarkTest02549Controller{})
}

func (c *BenchmarkTest02549Controller) Get() {
	c.handleRequest()
}

func (c *BenchmarkTest02549Controller) Post() {
	c.handleRequest()
}

func (c *BenchmarkTest02549Controller) handleRequest() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest02549="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamIndex(queryString, paramval)
	}
	if paramLoc == -1 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Ctx.Output.Body([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02549' in query string."))
		return
	}

	param := extractParamValue(queryString, paramval, paramLoc)

	bar := doSomething(param)
	c.handleDatabase(bar)
}

func findParamIndex(queryString, paramval string) int {
	return -1 // Implement logic to find parameter index
}

func extractParamValue(queryString, paramval string, paramLoc int) string {
	return "" // Implement logic to extract parameter value
}

func doSomething(param string) string {
	return param + "_SafeStuff"
}

func (c *BenchmarkTest02549Controller) handleDatabase(bar string) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	c.Ctx.Output.Body(output)
}
