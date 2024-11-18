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

type BenchmarkTest00350 struct {
	web.Controller
}

func (c *BenchmarkTest00350) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00350) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00350) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00350")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map51742 := make(map[string]interface{})
	map51742["keyA-51742"] = "a_Value"
	map51742["keyB-51742"] = param
	map51742["keyC"] = "another_Value"
	bar = map51742["keyB-51742"].(string)
	bar = map51742["keyA-51742"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func main() {
	web.Router("/crypto-00/BenchmarkTest00350", &BenchmarkTest00350{})
	web.Run()
}
