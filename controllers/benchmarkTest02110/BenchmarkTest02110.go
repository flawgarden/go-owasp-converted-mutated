package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest02110 struct{}

func (b *BenchmarkTest02110) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.FormValue("BenchmarkTest02110")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)
	fileName := fmt.Sprintf("testfiles/%s", bar)

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file:", fileName)
		return
	}
	defer fos.Close()

	w.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", fileName)))
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02110", &BenchmarkTest02110{})
	http.ListenAndServe(":8080", nil)
}
