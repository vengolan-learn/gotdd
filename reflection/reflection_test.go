package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         *Person
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			&Person{
				"Rama",
				Profile{22, "Sakethnagar"},
			},
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
