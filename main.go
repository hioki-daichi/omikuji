package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var nowFunc = time.Now

func main() {
	rand.Seed(nowFunc().UnixNano())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var result string

	if isDuringTheNewYear() {
		result = allFortunes()[0]
	} else {
		result = drawFortune()
	}

	fmt.Fprint(w, result)
}

func drawFortune() string {
	ss := allFortunes()
	return ss[rand.Intn(len(ss))]
}

func allFortunes() []string {
	return []string{"大吉", "中吉", "小吉", "吉", "末吉", "凶", "大凶"}
}

func isDuringTheNewYear() bool {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	_, month, day := nowFunc().In(loc).Date()
	if month == time.January && (day == 1 || day == 2 || day == 3) {
		return true
	}
	return false
}
