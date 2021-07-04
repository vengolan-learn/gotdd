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
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
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

	//Separate function to check map output since the order is not guaranteed...
	t.Run("with maps", func(t *testing.T) {

		aMap := map[string]string{
			"Foo": "FooVal",
			"Bar": "BarVal",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "FooVal")
		assertContains(t, got, "BarVal")

	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Moscow"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Moscow"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, but got %v", want, got)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q, but didn't", haystack, needle)
	}

}
