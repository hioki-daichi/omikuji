package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/hioki-daichi/omikuji-server/datehelper"
	"github.com/hioki-daichi/omikuji-server/fortune"
	"github.com/hioki-daichi/omikuji-server/jsonhelper"
	"github.com/hioki-daichi/omikuji-server/person"
)

var nowFunc = time.Now
var isDuringTheNewYearFunc = datehelper.IsDuringTheNewYear

func init() {
	rand.Seed(nowFunc().UnixNano())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var result fortune.Fortune
	if isDuringTheNewYearFunc() {
		result = fortune.Daikichi
	} else {
		result = fortune.DrawFortune()
	}

	p := person.NewPerson(result)

	json, err := jsonhelper.ToJSON(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, json)
}
