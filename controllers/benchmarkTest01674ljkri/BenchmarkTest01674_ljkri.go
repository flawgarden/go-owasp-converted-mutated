package controllers

import (
"fmt"
"net/http"
"net/url"
"os/exec"
)

type BenchmarkTest01674 struct{}

func (b *BenchmarkTest01674) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery

stringReturner := func() func() string {
    return func() string {
        return queryString
    }
}

stringRet := stringReturner()
queryString = stringRet()

	paramval := "BenchmarkTest01674="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		for i := 0; i < len(queryString)-len(paramval); i++ {
			if queryString[i:i+len(paramval)] == paramval {
				paramLoc = i
				break
			}
		}
	}

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01674"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)
	for i := paramLoc + len(paramval); i < len(queryString); i++ {
		if queryString[i] == '&' {
			ampersandLoc = i
			break
		}
	}
	if ampersandLoc != len(queryString) {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	a1 := "sh"
	a2 := "-c"
	args := []string{a1, a2, "echo " + bar}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Problem executing cmdi", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	thing := createThing()
	bar := thing.doSomething(param)
	return bar
}

type ThingInterface interface {
	doSomething(string) string
}

func createThing() ThingInterface {
	return &thing{}
}

type thing struct{}

func (t *thing) doSomething(param string) string {
	return param
}

func main() {
	http.Handle("/cmdi-01/BenchmarkTest01674", &BenchmarkTest01674{})
	http.ListenAndServe(":8080", nil)
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


