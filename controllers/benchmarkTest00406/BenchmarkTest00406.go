package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
)

type BenchmarkTest00406Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest00406Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	c.doPost()
}

func (c *BenchmarkTest00406Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Request.URL.Query().Get("BenchmarkTest00406")
	if param == "" {
		param = ""
	}

	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	var cmd *exec.Cmd
	if isWindows() {
		cmd = exec.Command("cmd.exe", "/c", "echo "+bar)
	} else {
		cmd = exec.Command("sh", "-c", "ping -c1 "+bar)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		http.Error(c.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	c.ResponseWriter.Write(output)
}

func isWindows() bool {
	return exec.Command("cmd", "/c", "ver").Run() == nil
}
