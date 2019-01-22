package hcl2shim

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestValuesSDKEquivalent(t *testing.T) {
	tests := []struct {
		A, B cty.Value
		Want bool
	}{
		// Strings
		{
			cty.StringVal("hello"),
			cty.StringVal("hello"),
			true,
		},
		{
			cty.StringVal("hello"),
			cty.StringVal("world"),
			false,
		},
		{
			cty.StringVal("hello"),
			cty.StringVal(""),
			false,
		},
		{
			cty.NullVal(cty.String),
			cty.StringVal(""),
			true,
		},

		// Numbers
		{
			cty.NumberIntVal(1),
			cty.NumberIntVal(1),
			true,
		},
		{
			cty.NumberIntVal(1),
			cty.NumberIntVal(2),
			false,
		},
		{
			cty.NumberIntVal(1),
			cty.Zero,
			false,
		},
		{
			cty.NullVal(cty.Number),
			cty.Zero,
			true,
		},

		// Bools
		{
			cty.True,
			cty.True,
			true,
		},
		{
			cty.True,
			cty.False,
			false,
		},
		{
			cty.NullVal(cty.Bool),
			cty.False,
			true,
		},

		// Mixed primitives
		{
			cty.StringVal("hello"),
			cty.False,
			false,
		},
		{
			cty.StringVal(""),
			cty.False,
			true,
		},
		{
			cty.NumberIntVal(0),
			cty.False,
			true,
		},
		{
			cty.StringVal(""),
			cty.NumberIntVal(0),
			true,
		},
		{
			cty.NullVal(cty.Bool),
			cty.NullVal(cty.Number),
			true,
		},
		{
			cty.StringVal(""),
			cty.NullVal(cty.Number),
			true,
		},

		// Lists
		{
			cty.ListValEmpty(cty.String),
			cty.NullVal(cty.List(cty.String)),
			true,
		},
	}

	run := func(t *testing.T, a, b cty.Value, want bool) {
		got := ValuesSDKEquivalent(a, b)

		if got != want {
			t.Errorf("wrong result\nfor: %#v ≈ %#v\ngot %#v, but want %#v", a, b, got, want)
		}
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%#v ≈ %#v", test.A, test.B), func(t *testing.T) {
			run(t, test.A, test.B, test.Want)
		})
		// This function is symmetrical, so we'll also test in reverse so
		// we don't need to manually copy all of the test cases. (But this does
		// mean that one failure normally becomes two, of course!)
		t.Run(fmt.Sprintf("%#v ≈ %#v", test.B, test.A), func(t *testing.T) {
			run(t, test.B, test.A, test.Want)
		})
	}
}
