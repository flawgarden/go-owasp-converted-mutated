package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00368 struct {
	web.Controller
}

func (c *BenchmarkTest00368) Get() {
	c.doPost()
}

func (c *BenchmarkTest00368) Post() {
	c.doPost()
}

func (c *BenchmarkTest00368) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00368")
	if param == "" {
		param = ""
	}

	bar := param // Escape HTML if necessary

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id, err := strconv.Atoi(bar)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Invalid ID", http.StatusBadRequest)
		return
	}

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%d", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "User not found", http.StatusNotFound)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error marshaling user", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
