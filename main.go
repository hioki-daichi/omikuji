package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, drawFortune())
}

func drawFortune() string {
	ss := []string{"大吉", "中吉", "小吉", "吉", "末吉", "凶", "大凶"}
	return ss[rand.Intn(len(ss))]
}
