package controllers

import (
"database/sql"
"encoding/json"
"fmt"
"go-sec-code/models"
"github.com/beego/beego/v2/server/web"
_ "github.com/go-sql-driver/mysql"
"sync"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02641Controller struct {
	web.Controller
}

func (c *BenchmarkTest02641Controller) Get() {
	c.handleRequest()
}

func (c *BenchmarkTest02641Controller) Post() {
	c.handleRequest()
}

func (c *BenchmarkTest02641Controller) handleRequest() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := c.Ctx.Input.URI()

dataChannel := make(chan Data, 1)
dataChannel <- Data{Value: queryString}

var wg sync.WaitGroup
wg.Add(1)

go func() {
    data := <-dataChannel
    data.Value = "constant_string"
    dataChannel <- data
}()

wg.Wait()

readData := <-dataChannel
queryString = readData.Value

	paramVal := "BenchmarkTest02641="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramVal)
	}

	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02641")))
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from user where username='foo' and password='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user models.User
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("No results returned for query: %s", sqlStr)))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	return param
}

func indexOf(s, substr string) int {
	return -1 // Implement logic to return the index of substr in s
}

func main() {
	web.Run()
}
