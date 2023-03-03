package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testwebservermod/config"
	"testwebservermod/utils"
)

func Test_server_addroutes(t *testing.T) {

	// Initialize
	s := NewServer(&config.Config{Addr: ":8080", TemplatePath: "../templates/"})
	mux, _ := s.AddRoutes()
	tests := []struct {
		Name           string
		Route          string
		Method         string
		Body           io.Reader
		ExpectedStatus int
		ExpectedOutput interface{}
	}{
		{"ping test", "/ping", "GET", nil, 200, &utils.ApiError{Code: 200, Message: "pong"}},
		{"home test", "/home", "GET", nil, 200, "Home page"},
	}
	// Invoke
	for _, test := range tests {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(test.Method, test.Route, test.Body)
		mux.ServeHTTP(rec, req)

		// Verify
		if test.Route == "/ping" {
			v := &utils.ApiError{}
			json.NewDecoder(rec.Body).Decode(v)
			op := &utils.ApiError{Code: 200, Message: "pong"}
			if v.Code != op.Code && v.Message != op.Message {
				t.Fail()
			}
		}
		if test.Route == "/home" {
			if !strings.Contains(string((rec.Body.Bytes())), test.ExpectedOutput.(string)) || test.ExpectedStatus != rec.Code {
				t.Fail()
			}
		}

	}

}
