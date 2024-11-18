package controllers

import (
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest02379Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02379Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	c.doPost()
}

func (c *BenchmarkTest02379Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Request.URL.Query().Get("BenchmarkTest02379")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	startURIslashes := ""
	if os.PathSeparator == '\\' {
		startURIslashes = "/"
	} else {
		startURIslashes = "//"
	}

	fileURI := filepath.Join(startURIslashes, "testfiles", bar)
	fileTarget := &os.File{}
	var err error
	if fileTarget, err = os.Open(fileURI); err != nil {
		c.ResponseWriter.Write([]byte("Error: " + err.Error()))
		return
	}
	defer fileTarget.Close()

	c.ResponseWriter.Write([]byte("Access to file: '" + fileTarget.Name() + "' created.<br>"))
	if _, err := os.Stat(fileTarget.Name()); err == nil {
		c.ResponseWriter.Write([]byte("And file already exists.<br>"))
	} else {
		c.ResponseWriter.Write([]byte("But file doesn't exist yet.<br>"))
	}
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // Remove the 1st safe value
		bar = valuesList[1]                                    // Get the last 'safe' value
	}
	return bar
}
