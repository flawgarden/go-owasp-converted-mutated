package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type BenchmarkTest00643 struct{}

func (h *BenchmarkTest00643) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00643")

map787234 := make(map[string]string)
map787234["ahWTo"] = "tkGxg"
map787234["tkGxg"] = param
map787234 = make(map[string]string)
param = map787234["tkGxg"]

	if param == "" {
		param = ""
	}

	sbxyz31207 := strings.Builder{}
	sbxyz31207.WriteString(param)
	bar := sbxyz31207.String() + "_SafeStuff"

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func TestBenchmarkTest00643(t *testing.T) {
	handler := &BenchmarkTest00643{}
	req := httptest.NewRequest(http.MethodPost, "/xss-01/BenchmarkTest00643", nil)
	req.Form.Add("BenchmarkTest00643", "test")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	expected := "test_SafeStuff"
	if rr.Body.String() != expected {
		t.Errorf("expected %v, got %v", expected, rr.Body.String())
	}
}
