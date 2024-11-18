package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01781Controller struct {
	web.Controller
}

func (c *BenchmarkTest01781Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01781Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01781Controller) DoPost() {
	response := c.Ctx.ResponseWriter
	request := c.Ctx.Request
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	idStr := request.FormValue("BenchmarkTest01781")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "Invalid id", http.StatusBadRequest)
		return
	}

	user, err := getUserByID(id)
	if err != nil {
		http.Error(response, "User not found", http.StatusNotFound)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(response, "Error marshaling user", http.StatusInternalServerError)
		return
	}

	response.Write(output)
}

func getUserByID(id int) (models.User, error) {
	var user models.User
	source := "root:password@tcp(127.0.0.1:3306)/goseccode"
	db, err := sql.Open("mysql", source)
	if err != nil {
		return user, err
	}
	defer db.Close()

	sqlStr := "SELECT * FROM user WHERE id = ?"
	err = db.QueryRow(sqlStr, id).Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}
