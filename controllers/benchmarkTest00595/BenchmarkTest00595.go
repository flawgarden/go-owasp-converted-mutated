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

type BenchmarkTest00595Controller struct {
	web.Controller
}

func (c *BenchmarkTest00595Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00595Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Input.Params()
	for name := range names {
		values := c.GetStrings(name)
		if values != nil {
			for _, value := range values {
				if value == "BenchmarkTest00595" {
					param = name
					flag = false
				}
			}
		}
		if !flag {
			break
		}
	}

	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("foo")
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}

	var users []models.User
	err = stmt.QueryRow().Scan(&users)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}

	output, err := json.Marshal(users)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	c.Ctx.Output.Body(output)
}
