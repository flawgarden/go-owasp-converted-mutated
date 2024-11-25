//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Gosec analysis results: [89, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02534/BenchmarkTest02534.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_struct_positive 
//Used extensions: 
//Program:
package controllers

import (
"database/sql"
"encoding/json"
"fmt"
"net/http"
"go-sec-code/models"
"github.com/beego/beego/v2/server/web"
_ "github.com/go-sql-driver/mysql"
"sync"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02534 struct {
	web.Controller
}

func (c *BenchmarkTest02534) Get() {
	c.Post()
}

func (c *BenchmarkTest02534) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.Ctx.Request.URL.Query()["BenchmarkTest02534"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

dataChannel := make(chan Data, 1)
dataChannel <- Data{Value: bar}

var wg sync.WaitGroup
wg.Add(1)

go func() {
    data := <-dataChannel
    data.Value += "suffix"
    dataChannel <- data
}()

wg.Wait()

readData := <-dataChannel
bar = readData.Value

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	conn, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	stmt, err := conn.Prepare(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query("foo")
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	return param
}
