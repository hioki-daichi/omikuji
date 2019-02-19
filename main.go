package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/hioki-daichi/omikuji-server/internal/datehelper"
	"github.com/hioki-daichi/omikuji-server/internal/form"
	"github.com/hioki-daichi/omikuji-server/internal/fortune"
	"github.com/hioki-daichi/omikuji-server/internal/jsonhelper"
)

// for testing
var nowFunc = time.Now
var isDuringTheNewYearFunc = datehelper.IsDuringTheNewYear
var toJSONFunc = jsonhelper.ToJSON

func init() {
	rand.Seed(nowFunc().UnixNano())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	var ftn fortune.Fortune
	if isDuringTheNewYearFunc() {
		ftn = fortune.Daikichi
	} else {
		ftn = fortune.DrawFortune()
	}

	p := form.NewRootForm(r).NewPerson(ftn)

	p.Validate()

	var v interface{}
	if len(p.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		v = map[string][]string{"errors": p.Errors}
	} else {
		v = p
	}

	json, err := toJSONFunc(v)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	fmt.Fprint(w, json)
}
