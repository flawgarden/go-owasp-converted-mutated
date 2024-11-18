package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
)

func BenchmarkTest00361(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00361")
	if param == "" {
		param = ""
	}

	var bar string
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

	fileName := "/path/to/test/files/" + bar

	fis, err := os.Open(fileName)
	if err != nil {
		http.Error(w, "Couldn't open FileInputStream on file: '"+fileName+"'", http.StatusInternalServerError)
		return
	}
	defer fis.Close()

	b, err := ioutil.ReadAll(fis)
	if err != nil {
		http.Error(w, "Couldn't read file: '"+fileName+"'", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n"))
	w.Write(b)
}

func main() {
	http.HandleFunc("/pathtraver-00/BenchmarkTest00361", BenchmarkTest00361)
	http.ListenAndServe(":8080", nil)
}
