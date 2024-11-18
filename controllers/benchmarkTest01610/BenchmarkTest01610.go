package controllers

import (
	"database/sql"
	"os"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01610Controller struct {
	web.Controller
}

func (c *BenchmarkTest01610Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01610Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest01610")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := new(Test).doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cmd := ""
	if strings.Contains(strings.ToLower(os.Getenv("OS")), "windows") {
		cmd = "cmd /C echo "
	}

	r := exec.Command(cmd + bar)

	p, err := r.Output()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Problem executing cmd - TestCase: " + err.Error()))
		return
	}
	c.Ctx.ResponseWriter.Write(p)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = string([]byte(param)) // Здесь необязательно использование кодирования Base64, можно заменить по необходимости
	}
	return bar
}
