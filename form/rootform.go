package form

import (
	"net/http"

	"github.com/hioki-daichi/omikuji-server/fortune"
	"github.com/hioki-daichi/omikuji-server/person"
)

// RootForm has name.
type RootForm struct {
	name string
}

// NewRootForm returns a form for the route of "/".
func NewRootForm(req *http.Request) *RootForm {
	f := &RootForm{}
	nameParam := req.URL.Query().Get("name")
	if nameParam != "" {
		f.name = nameParam
	} else {
		f.name = "Gopher"
	}
	return f
}

// NewPerson generates a person according to the content of form.
func (f *RootForm) NewPerson(ftn fortune.Fortune) *person.Person {
	return person.NewPerson(f.name, ftn)
}
