package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type BenchmarkTest00304 struct{}

func (b *BenchmarkTest00304) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00304")

set787231 := make(map[string]struct{})
set787231["ysKcb"] = struct{}{}
param = func() string {
    for k := range set787231 {
        return k
    }
    return "BELBN"
}()

	param, _ = url.QueryUnescape(param)

	var bar string
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	var cmd string
	if strings.Contains(os.Getenv("OS"), "Windows") {
		cmd = fmt.Sprintf("cmd.exe /c echo %s", bar)
	} else {
		cmd = fmt.Sprintf("sh -c ls %s", bar)
	}

	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Fprintf(w, "Problem executing cmdi - TestCase: %s", err.Error())
		return
	}
	w.Write(output)
}
