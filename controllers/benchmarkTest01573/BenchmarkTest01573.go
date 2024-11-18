package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest01573Controller struct{}

func (c *BenchmarkTest01573Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	values := r.URL.Query()["BenchmarkTest01573"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := c.doSomething(param)

	fileName := fmt.Sprintf("%s/%s", os.Getenv("TESTFILES_DIR"), bar)
	fis, err := os.Open(fileName)
	if err != nil {
		http.Error(w, "Problem getting FileInputStream: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	w.Write([]byte(fmt.Sprintf("The beginning of file: '%s' is:\n\n%s", bar, string(b[:size]))))
}

func (c *BenchmarkTest01573Controller) doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}
