package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var nowFunc = time.Now

type fortune string

const (
	daikichi fortune = "大吉"
	chukichi fortune = "中吉"
	shokichi fortune = "小吉"
	kichi    fortune = "吉"
	suekichi fortune = "末吉"
	kyo      fortune = "凶"
	daikyo   fortune = "大凶"
)

type person struct {
	Fortune fortune `json:"fortune"`
}

func main() {
	rand.Seed(nowFunc().UnixNano())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var result fortune

	if isDuringTheNewYear() {
		result = daikichi
	} else {
		result = drawFortune()
	}

	p := newPerson(result)
	json, err := generateJSON(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, json)
}

func newPerson(f fortune) *person {
	return &person{Fortune: f}
}

func generateJSON(p *person) (string, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func drawFortune() fortune {
	fs := allFortunes()
	return fs[rand.Intn(len(fs))]
}

func allFortunes() []fortune {
	return []fortune{daikichi, chukichi, shokichi, kichi, suekichi, kyo, daikyo}
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
