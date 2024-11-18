package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01912 struct {
	web.Controller
}

func (c *BenchmarkTest01912) Get() {
	c.Post()
}

func (c *BenchmarkTest01912) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Input.Header("BenchmarkTest01912")
	param = strings.TrimSpace(param)
	bar := doSomething(param)

	// Hashing and file writing logic would go here
	// For the purpose of this example, we'll just respond with the hashed value

	output, err := json.Marshal(map[string]string{"message": "Value processed", "value": bar})
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	c.Ctx.Output.Body(output)
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
