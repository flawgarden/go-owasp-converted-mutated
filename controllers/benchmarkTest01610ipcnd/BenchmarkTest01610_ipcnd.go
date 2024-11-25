//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01610/BenchmarkTest01610.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/assertions.tmt with name type_assertion_with_struct_pointer_negative 
//Used extensions: 
//Program:
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

var i123 interface{} = &Anon{Value1: bar}
if ptr, ok := i123.(*EmbeddedStruct); ok {
     bar = ptr.Field1
} else {
    bar = "dbNio"
}

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
