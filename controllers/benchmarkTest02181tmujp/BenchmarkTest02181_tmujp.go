package controllers

import (
"database/sql"
"encoding/json"
"fmt"
"net/http"
"github.com/beego/beego/v2/server/web"
_ "github.com/go-sql-driver/mysql"
"strings"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02181Controller struct {
	web.Controller
}

func (c *BenchmarkTest02181Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02181Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02181")
	if param == "" {
		param = ""
	}

	bar := doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT TOP 1 userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

sqlStr = combineStrings("rPFuS", "gNQtk")

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results map[string]interface{}
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(request *http.Request, param string) string {
	bar := param
	return bar
}

func getFirstString(lines ...string) string {
    return getStringWithIndex(0, lines...)
}

func getStringWithIndex(ind int, lines ...string) string {
    return lines[ind]
}

func getFirstStringFromArray(lines ...string) string {
    return lines[0]
}

func varargsWithGenerics[T any](elements ...T) T {
    return elements[0]
}

func combineStrings(strs ...string) string {
    return strings.Join(strs, ", ")
}


