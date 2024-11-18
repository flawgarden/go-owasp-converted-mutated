package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00605Controller struct {
	web.Controller
}

func (c *BenchmarkTest00605Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00605Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Input.Params()

	for name := range names {
		values := c.GetStrings(name)
		if values != nil {
			for _, value := range values {
				if value == "BenchmarkTest00605" {
					param = name
					flag = false
					break
				}
			}
		}
		if !flag {
			break
		}
	}

	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output := map[string]string{"message": "Update complete"}
	jsonOutput, _ := json.Marshal(output)
	c.Ctx.ResponseWriter.Write(jsonOutput)
}
