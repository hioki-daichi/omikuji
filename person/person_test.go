package person

import (
	"reflect"
	"testing"

	"github.com/hioki-daichi/omikuji-server/fortune"
)

func TestPerson_NewPerson(t *testing.T) {
	expected := "*person.Person"

	p := NewPerson("Gopher", fortune.Daikichi)

	actual := reflect.TypeOf(p).String()
	if actual != expected {
		t.Errorf(`unexpected : expected: "%s" actual: "%s"`, expected, actual)
	}
}
