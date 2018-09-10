package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain_handler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	expectedStatusCode := http.StatusOK
	actualStatusCode := rw.StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf(`unexpected status code: expected: "%d" actual: "%d"`, expectedStatusCode, actualStatusCode)
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatalf("err %s", err)
	}

	expectedResponseBody := "大吉"
	actualResponseBody := string(b)
	if actualResponseBody != expectedResponseBody {
		t.Errorf(`unexpected response body: expected: "%s" actual: "%s"`, expectedResponseBody, actualResponseBody)
	}
}
