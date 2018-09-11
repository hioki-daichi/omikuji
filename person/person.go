/*
Package person is a package that manages processing around person.
*/
package person

import (
	"github.com/hioki-daichi/omikuji-server/fortune"
)

// Person has Name and Fortune.
type Person struct {
	Name    string          `json:"name"`
	Fortune fortune.Fortune `json:"fortune"`
}

// NewPerson generates a new person.
func NewPerson(n string, f fortune.Fortune) *Person {
	return &Person{
		Name:    n,
		Fortune: f,
	}
}
