package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ThingInterface interface {
	DoSomething(param string) string
}

type ThingFactory struct{}

func (tf *ThingFactory) CreateThing() ThingInterface {
	return &Thing{}
}

type Thing struct{}

func (t *Thing) DoSomething(param string) string {
	return fmt.Sprintf("output_based_on_%s", param)
}

func BenchmarkTest00527(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	flag := true
	for name, values := range r.URL.Query() {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest00527" {
				param = name
				flag = false
				break
			}
		}
	}

	factory := &ThingFactory{}
	thing := factory.CreateThing()
	bar := thing.DoSomething(param)

	fileTarget := fmt.Sprintf("%s/Test.txt", bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", htmlEscape(fileTarget))))

	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func htmlEscape(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&", "&amp;"), "<", "&lt;")
}

func main() {
	http.HandleFunc("/pathtraver-00/BenchmarkTest00527", BenchmarkTest00527)
	http.ListenAndServe(":8080", nil)
}
