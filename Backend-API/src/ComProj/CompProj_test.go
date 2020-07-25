// main_test.go
package main

import (
	"os"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestGetNonProfits(t *testing.T) {
    req, _ := http.NewRequest("GET", "/NonProfits", nil)
    response := executeRequest(req)
    fmt.Println(response)
    //var m map[string]interface{}
    //json.Unmarshal(response.Body.Bytes(), &m)

    checkResponseCode(t, http.StatusOK, response.Code)
}

// func TestCreateNonProfits(t *testing.T) {
//     req, _ := http.NewRequest("POST", "/NonProfits", nil)
//     response := executeRequest(req)
//     fmt.Println(response)
//     //var m map[string]interface{}
//     //json.Unmarshal(response.Body.Bytes(), &m)

//     checkResponseCode(t, http.StatusOK, response.Code)
// }

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

