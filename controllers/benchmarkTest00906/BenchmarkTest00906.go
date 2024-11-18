package controllers

import (
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00906Controller struct {
	web.Controller
}

func (c *BenchmarkTest00906Controller) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest00906Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00906")

	bar := ""
	if param != "" {
		encoded := []byte(param) // тут используется простая кодировка для примера
		bar = string(encoded)
	}

	var cmd string
	osName := "Linux" // Замените на логику для определения ОС

	if osName == "Windows" {
		cmd = fmt.Sprintf("cmd.exe /c echo %s", bar)
	} else {
		cmd = fmt.Sprintf("sh -c ping -c1 %s", bar)
	}

	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Problem executing cmdi - TestCase: %s", err.Error()))
		return
	}
	c.Ctx.WriteString(string(output))
}
