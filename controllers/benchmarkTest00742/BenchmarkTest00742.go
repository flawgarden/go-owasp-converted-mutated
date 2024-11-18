package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
)

type BenchmarkTest00742 struct{}

func (b *BenchmarkTest00742) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00742")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map62435 := make(map[string]interface{})
	map62435["keyA-62435"] = "a_Value"
	map62435["keyB-62435"] = param
	map62435["keyC"] = "another_Value"
	bar = map62435["keyB-62435"].(string)
	bar = map62435["keyA-62435"].(string)

	var cmd string
	if isWindows() {
		cmd = "cmd /C echo " + bar
	}

	argsEnv := []string{"Foo=bar"}
	cmdExec := exec.Command("sh", "-c", cmd)
	cmdExec.Env = argsEnv
	output, err := cmdExec.CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func isWindows() bool {
	return exec.Command("uname").Run() != nil
}

func main() {
	http.Handle("/cmdi-00/BenchmarkTest00742", &BenchmarkTest00742{})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
