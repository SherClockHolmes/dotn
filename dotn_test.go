package dotn

import (
	"testing"
)

type Nested struct {
	Slice []string `json:"slice"`
}

type Root struct {
	Nested map[string]*Nested `json:"nested"`
}

func TestSet(t *testing.T) {
	dot := "nested.key1.slice.0"
	value := "three"

	root := &Root{
		Nested: map[string]*Nested{
			"key1": &Nested{
				Slice: []string{"one", "two"},
			},
		},
	}

	// Set the value via the dot string
	if err := Set(dot, value, root); err != nil {
		t.Fatal(err)
	}

	if root.Nested["key1"].Slice[0] != value {
		t.Fatalf(
			"Incorreect status code, expected=%s, got=%s",
			value,
			root.Nested["n1"].Slice[0],
		)
	}
}
