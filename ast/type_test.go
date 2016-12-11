package ast

import (
	"fmt"
	"testing"
)

func TestTypeCompare(t *testing.T) {
	// This test is intended to ensure that the Go types modelling the
	// HIL types are never extended in such a way as to make equality
	// testing not work properly.
	tests := []struct {
		TypeA Type
		TypeB Type
		Equal bool
	}{
		{
			TypeInt,
			TypeInt,
			true,
		},
		{
			TypeInt,
			TypeInvalid,
			false,
		},
		{
			TypeList{TypeInt},
			TypeList{TypeInt},
			true,
		},
		{
			TypeList{TypeInt},
			TypeList{TypeString},
			false,
		},
		{
			TypeMap{TypeInt},
			TypeMap{TypeInt},
			true,
		},
		{
			TypeMap{TypeInt},
			TypeMap{TypeString},
			false,
		},
		{
			TypeList{TypeInt},
			TypeMap{TypeInt},
			false,
		},
	}

	for _, test := range tests {
		var testName string
		if test.Equal {
			testName = fmt.Sprintf("%#v==%#v", test.TypeA, test.TypeB)
		} else {
			testName = fmt.Sprintf("%#v!=%#v", test.TypeA, test.TypeB)
		}
		t.Run(testName, func(t *testing.T) {
			got := test.TypeA == test.TypeB
			if want := test.Equal; got != want {
				t.Errorf("got %#v; want %#v", got, want)
			}
		})
	}
}

func TestTypePrintable(t *testing.T) {
	// The primary reason for this test is actually to ensure that
	// all of the type references below are compilable. The testing
	// of the "Printable" method is secondary and rather pointless for
	// most of them, but marginally interesting for the collection types.
	tests := []struct {
		Type     Type
		Expected string
	}{
		{
			TypeInvalid,
			"invalid type",
		},
		{
			TypeAny,
			"any type",
		},
		{
			TypeBool,
			"bool",
		},
		{
			TypeString,
			"string",
		},
		{
			TypeInt,
			"int",
		},
		{
			TypeFloat,
			"float",
		},
		{
			TypeUnknown,
			"unknown",
		},
		{
			TypeList{TypeString},
			"list of string",
		},
		{
			TypeList{TypeList{TypeString}},
			"list of list of string",
		},
		{
			TypeMap{TypeString},
			"map of string",
		},
		{
			TypeMap{TypeMap{TypeString}},
			"map of map of string",
		},
	}

	for _, test := range tests {
		t.Run(test.Expected, func(t *testing.T) {
			got := test.Type.Printable()
			if want := test.Expected; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}

func TestTypeGoString(t *testing.T) {
	tests := []struct {
		Type     Type
		Expected string
	}{
		{
			TypeInvalid,
			"ast.TypeInvalid",
		},
		{
			TypeAny,
			"ast.TypeAny",
		},
		{
			TypeBool,
			"ast.TypeBool",
		},
		{
			TypeString,
			"ast.TypeString",
		},
		{
			TypeInt,
			"ast.TypeInt",
		},
		{
			TypeFloat,
			"ast.TypeFloat",
		},
		{
			TypeUnknown,
			"ast.TypeUnknown",
		},
		{
			TypeList{TypeString},
			"ast.TypeList{ast.TypeString}",
		},
		{
			TypeList{TypeList{TypeString}},
			"ast.TypeList{ast.TypeList{ast.TypeString}}",
		},
		{
			TypeMap{TypeString},
			"ast.TypeMap{ast.TypeString}",
		},
		{
			TypeMap{TypeMap{TypeString}},
			"ast.TypeMap{ast.TypeMap{ast.TypeString}}",
		},
	}

	for _, test := range tests {
		t.Run(test.Expected, func(t *testing.T) {
			got := fmt.Sprintf("%#v", test.Type)
			if want := test.Expected; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}
