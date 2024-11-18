package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
)

type BenchmarkTest01353 struct{}

func (bt *BenchmarkTest01353) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01353")

	bar := bt.doSomething(r, param)

	var argList []string
	if isWindows() {
		argList = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		argList = []string{"sh", "-c", "echo " + bar}
	}

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Problem executing command", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func (bt *BenchmarkTest01353) doSomething(r *http.Request, param string) string {
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

	return bar
}

func isWindows() bool {
	return false // Simplified for example; you can use logic to check OS
}

func main() {
	http.Handle("/cmdi-01/BenchmarkTest01353", &BenchmarkTest01353{})
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
