package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Rama"},
			[]string{"Rama"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Rama", "Sakethnagar"},
			[]string{"Rama", "Sakethnagar"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("want %v, but got %v", test.ExpectedCalls, got)
			}

		})
	}

}
