package roops

import (
	"fmt"
	"testing"
)

var dataProvider = []struct {
	Expected, Given string
	Converted       bool
}{
	{Converted: false}, // Empty source value
	{Converted: true, Expected: "hello", Given: "руддщ"},                 // Simple string
	{Converted: true, Expected: "Foo", Given: "Ащщ"},                     // Lower and uppercase
	{Converted: true, Expected: "mono 83", Given: "ьщтщ 83"},             // Spaces and numbers
	{Converted: false, Expected: "foo ащщ", Given: "foo ащщ"},            // Mixed encodings
	{Converted: true, Expected: "Hello, world!", Given: "Руддщб цщкдв!"}, // Hello world
}

func TestToLatin(t *testing.T) {
	for _, data := range dataProvider {
		t.Run(fmt.Sprintf("%s->%s", data.Given, data.Expected), func(t *testing.T) {
			res, success := ToLatin(data.Given)
			if success == data.Converted {
				// Asserting value
				if res != data.Expected {
					t.Errorf(`Expected "%s" but converted to "%s"`, data.Expected, res)
				}
			} else {
				if data.Converted {
					t.Errorf(`ToLatin does not convert "%s" as expected (function returns false)`, data.Given)
				} else {
					t.Errorf(`ToLatin converts "%s" but should not do it (function returns true)`, data.Given)
				}
			}
		})
	}
}
