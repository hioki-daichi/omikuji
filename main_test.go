package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMain_handler_StatusCode(t *testing.T) {
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

func TestMain_handler_DuringTheNewYear(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	nowFunc = func() time.Time {
		return time.Date(2019, time.January, 1, 0, 0, 0, 0, loc)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatalf("err %s", err)
	}

	expected := "大吉"
	actual := string(b)
	if actual != expected {
		t.Errorf(`unexpected response body: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func TestMain_isDuringTheNewYear(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	cases := map[string]struct {
		year     int
		month    time.Month
		day      int
		expected bool
	}{
		"2018-12-31": {year: 2018, month: time.December, day: 31, expected: false},
		"2019-01-01": {year: 2019, month: time.January, day: 1, expected: true},
		"2019-01-02": {year: 2019, month: time.January, day: 2, expected: true},
		"2019-01-03": {year: 2019, month: time.January, day: 3, expected: true},
		"2019-01-04": {year: 2019, month: time.January, day: 4, expected: false},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			nowFunc = func() time.Time {
				return time.Date(c.year, c.month, c.day, 0, 0, 0, 0, loc)
			}

			expected := c.expected
			actual := isDuringTheNewYear()
			if actual != expected {
				t.Errorf(`expected: "%t" actual: "%t"`, expected, actual)
			}
		})
	}
}
