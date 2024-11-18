package controllers

import (
	"os"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest01401Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01401Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01401Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01401" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := new(Test).DoSomething(param)

	// Crypto logic
	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error opening file"))
		return
	}
	defer fw.Close()

	_, err = fw.WriteString("secret_value=" + bar + "\n")
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error writing to file"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Sensitive value: '" + bar + "' encrypted and stored<br/>"))
}

type Test struct{}

func (t *Test) DoSomething(param string) string {
	bar := ""

	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
