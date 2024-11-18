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

func init() {
	web.Router("/crypto-02/BenchmarkTest01741", &BenchmarkTest01741Controller{})
}

type BenchmarkTest01741Controller struct {
	web.Controller
}

func (c *BenchmarkTest01741Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01741Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01741")
	bar := new(Test).doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
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

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz82048 := []rune(param)
		bar = string(append(sbxyz82048[:len(sbxyz82048)-1], 'Z'))
	}
	return bar
}

func main() {
	web.Run()
}
