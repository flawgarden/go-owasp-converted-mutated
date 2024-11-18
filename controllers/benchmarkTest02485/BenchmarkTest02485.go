package controllers

import (
	"fmt"
	"net/http"

	"xorm.io/xorm"
)

type BenchmarkTest02485 struct {
	engine *xorm.Engine
}

func (bt *BenchmarkTest02485) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := r.Form["BenchmarkTest02485"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := bt.doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintf(w, "Formatted like: %s and %s.", "a", bar)
}

func (bt *BenchmarkTest02485) doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:password@tcp(127.0.0.1:3306)/goseccode")
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	bt := &BenchmarkTest02485{engine: engine}
	http.Handle("/xss-04/BenchmarkTest02485", bt)
	http.ListenAndServe(":8080", nil)
}
