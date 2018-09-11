/*
Package person is a package that manages processing around person.
*/
package person

import (
	"github.com/hioki-daichi/omikuji-server/fortune"
)

// Person has Fortune.
type Person struct {
	Fortune fortune.Fortune `json:"fortune"`
}

// NewPerson generates a new person.
func NewPerson(f fortune.Fortune) *Person {
	return &Person{Fortune: f}
}
