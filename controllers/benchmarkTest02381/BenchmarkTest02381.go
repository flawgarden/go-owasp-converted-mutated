package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest02381Controller struct{}

func (b *BenchmarkTest02381Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02381")
	if param == "" {
		param = ""
	}

	bar := doSomething(r, param)

	fileName := fmt.Sprintf("testfiles/%s", bar)

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	output := fmt.Sprintf("Now ready to write to file: %s", htmlEscape(fileName))
	w.Write([]byte(output))
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map75774 := make(map[string]interface{})
	map75774["keyA-75774"] = "a_Value"
	map75774["keyB-75774"] = param
	map75774["keyC"] = "another_Value"
	bar = map75774["keyB-75774"].(string)
	bar = map75774["keyA-75774"].(string)
	return bar
}

func htmlEscape(s string) string {
	safe := jsonEscape(s)
	return safe
}

func jsonEscape(s string) string {
	return s // здесь должны быть реализованы дополнительные функции экранирования
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02381", &BenchmarkTest02381Controller{})
	http.ListenAndServe(":8080", nil)
}
