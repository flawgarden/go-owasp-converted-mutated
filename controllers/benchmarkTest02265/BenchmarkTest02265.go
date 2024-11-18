package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"go-sec-code/models"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Database initialization code can go here
}

type BenchmarkTest02265Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02265Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02265Controller) Post() {
	queryParam := c.GetString("BenchmarkTest02265")
	bar := doSomething(queryParam)

	sqlStr := "{call " + bar + "}"

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			c.Ctx.Output.Body([]byte("Error processing request."))
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	c.Ctx.Output.Body(output)
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the last 'safe' value
	}
	return bar
}
