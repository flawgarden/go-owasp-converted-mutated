package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00938Controller struct {
	web.Controller
}

func (c *BenchmarkTest00938Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00938Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00938")

	bar := ""
	if param != "" {
		bar = string([]byte(param)) // Эмулируем декодирование Base64 для примера
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	user := models.User{}
	err = statement.QueryRow().Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
