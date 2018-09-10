package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain_handler_StatusCode(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	expected := http.StatusOK
	actual := rw.StatusCode
	if actual != expected {
		t.Errorf(`unexpected status code: expected: "%d" actual: "%d"`, expected, actual)
	}
}

func TestMain_handler_ResponseBody(t *testing.T) {
	cases := map[string]struct {
		seed     int64
		expected string
	}{
		"KYOU":     {seed: 0, expected: "凶"},
		"DAIKYOU":  {seed: 1, expected: "大凶"},
		"SUEKICHI": {seed: 2, expected: "末吉"},
		"KICHI":    {seed: 3, expected: "吉"},
		"CHUKICHI": {seed: 4, expected: "中吉"},
		"SHOKICHI": {seed: 5, expected: "小吉"},
		"DAICHIKI": {seed: 9, expected: "大吉"},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			// Can not use t.Parallel() because of rand.Seed
			rand.Seed(c.seed)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			handler(w, r)
			rw := w.Result()
			defer rw.Body.Close()

			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatalf("err %s", err)
			}

			expected := c.expected
			actual := string(b)
			if actual != expected {
				t.Errorf(`unexpected response body: expected: "%s" actual: "%s"`, expected, actual)
			}
		})
	}
}
