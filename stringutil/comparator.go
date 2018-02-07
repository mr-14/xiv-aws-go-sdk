package stringutil

import (
	"encoding/json"
	"reflect"
)

// EqualJSON compares the JSON strings
func EqualJSON(s1, s2 string) bool {
	var o1, o2 interface{}
	var err error

	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false
	}

	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false
	}

	return reflect.DeepEqual(o1, o2)
}
