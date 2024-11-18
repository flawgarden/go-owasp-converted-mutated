package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os/exec"
	"testing"
)

type BenchmarkTest01064 struct{}

func (b *BenchmarkTest01064) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest01064") != "" {
		param = r.Header.Get("BenchmarkTest01064")
	}

	param, _ = url.QueryUnescape(param)

	bar := b.newTest().doSomething(r, param)

	cmd := ""
	if os := r.Header.Get("User-Agent"); os != "" {
		cmd = fmt.Sprintf("echo %s", bar)
	}

	_, err := exec.Command("cmd", "/C", cmd).Output()
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem executing cmdi - TestCase: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (b *BenchmarkTest01064) newTest() *Test {
	return &Test{}
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map72463 := make(map[string]interface{})
	map72463["keyA-72463"] = "a-Value"
	map72463["keyB-72463"] = param
	map72463["keyC"] = "another-Value"
	bar = map72463["keyB-72463"].(string)
	return bar
}

func TestBenchmark(t *testing.T) {
	req := httptest.NewRequest("GET", "/cmdi-01/BenchmarkTest01064", nil)
	req.Header.Set("BenchmarkTest01064", "testParam")

	rec := httptest.NewRecorder()
	b := &BenchmarkTest01064{}
	b.doPost(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", rec.Code)
	}
}
