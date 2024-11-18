package controllers

import (
	"net/http"
	"os/exec"
	"strings"
)

type BenchmarkTest02340Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02340Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	c.doPost()
}

func (c *BenchmarkTest02340Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Request.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02340" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := c.doSomething(param)

	var cmd string
	var a1 string
	var a2 string
	var args []string
	osName := "os.name" // Replace with a proper call to detect the OS.

	if strings.Contains(osName, "Windows") {
		a1 = "cmd.exe"
		a2 = "/c"
		cmd = "echo "
		args = []string{a1, a2, cmd + bar}
	} else {
		a1 = "sh"
		a2 = "-c"
		cmd = "ls " // Replace with proper command if needed.
		args = []string{a1, a2, cmd + bar}
	}

	argsEnv := []string{"foo=bar"}
	r := exec.Command(args[0], args[1:]...)
	r.Env = argsEnv

	output, err := r.CombinedOutput()
	if err != nil {
		c.ResponseWriter.Write([]byte("Problem executing cmdi - TestCase"))
		c.ResponseWriter.Write([]byte(err.Error()))
		return
	}
	c.ResponseWriter.Write(output)
}

func (c *BenchmarkTest02340Controller) doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
