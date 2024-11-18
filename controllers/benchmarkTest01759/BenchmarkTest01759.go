package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01759 struct {
	web.Controller
}

func (c *BenchmarkTest01759) Get() {
	c.Post()
}

func (c *BenchmarkTest01759) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01759")
	bar := escape(param)

	hashAndStoreSensitiveValue(bar, c.Ctx.ResponseWriter)
}

func escape(param string) string {
	return strings.ReplaceAll(param, "<", "&lt;")
}

func hashAndStoreSensitiveValue(bar string, response http.ResponseWriter) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	response.Write(output)

	response.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", bar)))
}
