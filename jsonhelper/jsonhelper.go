/*
Package jsonhelper is a collection of convenient functions for manipulating JSON.
*/
package jsonhelper

import (
	"bytes"
	"encoding/json"
)

// ToJSON converts v to JSON.
func ToJSON(v interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(v); err != nil {
		return "", err
	}
	return buf.String(), nil
}
