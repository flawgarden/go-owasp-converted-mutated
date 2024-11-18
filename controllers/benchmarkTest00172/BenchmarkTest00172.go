package controllers

import (
	"net/url"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00172Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00172Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00172Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest00172")
	if param != "" {
		decodedParam, _ := url.QueryUnescape(param)

		bar := "safe!"
		map59408 := map[string]interface{}{
			"keyA-59408": "a-Value",
			"keyB-59408": decodedParam,
			"keyC":       "another-Value",
		}
		bar = map59408["keyB-59408"].(string)

		cmd := "your-command-here" // replace with actual command
		args := []string{cmd}
		argsEnv := []string{bar}

		r := exec.Command(args[0], args[1:]...)
		r.Env = append(r.Env, argsEnv...)

		output, err := r.CombinedOutput()
		if err != nil {
			c.Ctx.ResponseWriter.Write([]byte(err.Error()))
			return
		}
		c.Ctx.ResponseWriter.Write(output)
	}
}
