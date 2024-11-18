package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00425Controller struct {
	web.Controller
}

func (c *BenchmarkTest00425Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00425Controller) Post() {
	response := c.Ctx.ResponseWriter
	request := c.Ctx.Request

	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := request.FormValue("BenchmarkTest00425")
	if param == "" {
		param = ""
	}

	bar := param + "_SafeStuff"

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO session_data (key, value) VALUES (?, ?)", "userid", bar)
	if err != nil {
		panic(err)
	}

	output := fmt.Sprintf("Item: 'userid' with value: '%s' saved in session.", bar)
	response.Write([]byte(output))
}

func main() {
	web.Router("/trustbound-00/BenchmarkTest00425", &BenchmarkTest00425Controller{})
	web.Run()
}
