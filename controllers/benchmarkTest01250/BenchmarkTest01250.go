package controllers

import (
	"fmt"
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

type BenchmarkTest01250Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01250Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01250Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01250")
	if param == "" {
		param = ""
	}

	bar := c.doSomething(param)

	hashValue, err := hashValue(bar)
	if err != nil {
		panic(err)
	}

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fw.Close()

	if _, err := fw.WriteString("hash_value=" + hashValue + "\n"); err != nil {
		panic(err)
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", htmlEncode(param))))
}

func (c *BenchmarkTest01250Controller) doSomething(param string) string {
	// Mocking ThingInterface and its method doSomething for demonstration
	return param
}

func hashValue(input string) (string, error) {
	// Simulate hashing logic, replace with actual implementation
	return "hashed_value_example", nil
}

func htmlEncode(value string) string {
	return value // Replace with actual HTML encoding implementation
}
