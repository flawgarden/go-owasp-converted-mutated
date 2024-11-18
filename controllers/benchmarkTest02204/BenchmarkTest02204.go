package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/pathtraver-02/BenchmarkTest02204", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		param := r.FormValue("BenchmarkTest02204")
		bar := doSomething(param)

		fileName := filepath.Join("TESTFILES_DIR", bar)

		fos, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
			return
		}
		defer fos.Close()

		w.Header().Set("Content-Type", "text/html;charset=UTF-8")
		w.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", fileName)))
	})

	http.ListenAndServe(":8080", nil)
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
