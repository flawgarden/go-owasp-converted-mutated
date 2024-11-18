package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00620 struct{}

func (b *BenchmarkTest00620) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00620")
	if param == "" {
		param = ""
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	fileTarget := fmt.Sprintf("%s/%s", os.Getenv("TESTFILES_DIR"), bar)
	fmt.Fprintf(w, "Access to file: '%s' created.", fileTarget)
	if _, err := os.Stat(fileTarget); err == nil {
		fmt.Fprintf(w, " And file already exists.")
	} else {
		fmt.Fprintf(w, " But file doesn't exist yet.")
	}
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00620", &BenchmarkTest00620{})
	http.ListenAndServe(":8080", nil)
}
