package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00761Controller struct {
	web.Controller
}

func (c *BenchmarkTest00761Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00761Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	idParam := c.GetString("BenchmarkTest00761")
	var id int
	if idParam != "" {
		var err error
		id, err = strconv.Atoi(idParam)
		if err != nil {
			id = 0
		}
	} else {
		id = 0
	}

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%d", id)
	user := models.User{}

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
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
