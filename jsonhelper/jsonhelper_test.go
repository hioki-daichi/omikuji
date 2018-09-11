package jsonhelper

import (
	"testing"
)

func TestJsonhelper_ToJSON(t *testing.T) {
	foo := struct {
		Bar string `json:"bar"`
		Baz int    `json:"baz"`
	}{
		Bar: "barbar",
		Baz: 1,
	}

	actual, err := ToJSON(foo)
	if err != nil {
		t.Fatalf("err %s", err)
	}
	expected := "{\"bar\":\"barbar\",\"baz\":1}\n"
	if actual != expected {
		t.Errorf(`unexpected : expected: "%s" actual: "%s"`, expected, actual)
	}
}
