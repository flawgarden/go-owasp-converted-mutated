package controllers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01608 struct {
	web.Controller
}

func (c *BenchmarkTest01608) Get() {
	c.Post()
}

func (c *BenchmarkTest01608) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.Ctx.Input.Query("BenchmarkTest01608")
	var param string
	if values != "" {
		param = values
	} else {
		param = ""
	}

	bar := new(Test).doSomething(param)

	cmd := getInsecureOSCommandString()
	args := []string{cmd}
	argsEnv := []string{bar}

	err := runCommand(args, argsEnv)
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map55872 := make(map[string]interface{})
	map55872["keyA-55872"] = "a_Value"
	map55872["keyB-55872"] = param
	map55872["keyC"] = "another_Value"
	bar = map55872["keyB-55872"].(string)
	bar = map55872["keyA-55872"].(string)
	return bar
}

func getInsecureOSCommandString() string {
	return "your_command_here" // Replace with your command logic
}

func runCommand(args []string, argsEnv []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = argsEnv
	cmd.Dir = os.Getenv("user.dir")
	return cmd.Start()
}

func main() {
	web.Router("/cmdi-01/BenchmarkTest01608", &BenchmarkTest01608{})
	web.Run()
}
