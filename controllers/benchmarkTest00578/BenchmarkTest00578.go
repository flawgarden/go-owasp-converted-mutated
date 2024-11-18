package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00578Controller struct {
	web.Controller
}

func (c *BenchmarkTest00578Controller) Get() {
	id := c.GetString("id")
	if id == "" {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
			return
		}
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *BenchmarkTest00578Controller) Post() {
	c.Get()
}
