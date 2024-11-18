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

type BenchmarkTest02405Controller struct {
	web.Controller
}

func (c *BenchmarkTest02405Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02405Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest02405")
	if param == "" {
		param = ""
	}
	bar := doSomething(param)
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", param)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return string(output)
}

func main() {
	web.Router("/xss-04/BenchmarkTest02405", &BenchmarkTest02405Controller{})
	web.Run()
}
