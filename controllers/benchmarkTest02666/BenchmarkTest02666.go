package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest02666Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02666Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	c.doPost()
}

func (c *BenchmarkTest02666Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Request.URL.Query().Get("BenchmarkTest02666")
	bar := c.doSomething(param)

	fileName := fmt.Sprintf("testfiles/%s", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileInputStream on file: '" + fileName + "'")
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	c.ResponseWriter.Write([]byte(fmt.Sprintf("The beginning of file: '%s' is:\n\n", fileName)))
	c.ResponseWriter.Write([]byte(fmt.Sprintf("%s", string(b[:size]))))
}

func (c *BenchmarkTest02666Controller) doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
