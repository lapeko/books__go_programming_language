package pkg

import (
	"reflect"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	e := newEncoder()
	eT := reflect.TypeOf(e)
	expectT := reflect.TypeOf(&encoder{})
	if eT != expectT {
		t.Errorf("Encoder onstance has a wrong type %q when %q expected", eT, expectT)
	}
}

func TestEncodeStruct(t *testing.T) {
	tests := []struct {
		testStruct any
		expected   string
	}{
		{
			struct{ field int }{field: 1},
			`{ field: 1 }`,
		},
		{
			struct {
				num       int
				friendIds []int
			}{4, []int{1, 2, 3, 4}},
			`{ num: 4, friendIds: { 1, 2, 3, 4 } }`,
		},
	}

	for _, tt := range tests {
		e := newEncoder()
		e.encodeStruct(reflect.ValueOf(tt.testStruct))
		res := e.String()
		if res != tt.expected {
			t.Errorf("encodeStruct(%v) = %q expected %q", tt.testStruct, res, tt.expected)
		}
	}
}

func TestEncodeArray(t *testing.T) {
	tests := []struct {
		testArray any
		expected  string
	}{
		{
			[...]string{"val1", "test"},
			`{ "val1", "test" }`,
		},
		{
			[...]int{1, 3, 2, 5},
			`{ 1, 3, 2, 5 }`,
		},
	}

	for _, tt := range tests {
		e := newEncoder()
		e.encodeIterable(reflect.ValueOf(tt.testArray))
		res := e.String()
		if res != tt.expected {
			t.Errorf("encodeIterable(%v) = %q expected %q", tt.testArray, res, tt.expected)
		}
	}
}

func TestEncodeSlice(t *testing.T) {
	tests := []struct {
		testSlice any
		expected  string
	}{
		{
			[]string{"val1", "test"},
			`{ "val1", "test" }`,
		},
		{
			[]int{1, 3, 2, 5},
			`{ 1, 3, 2, 5 }`,
		},
	}

	for _, tt := range tests {
		e := newEncoder()
		e.encodeIterable(reflect.ValueOf(tt.testSlice))
		res := e.String()
		if res != tt.expected {
			t.Errorf("encodeIterable(%v) = %q expected %q", tt.testSlice, res, tt.expected)
		}
	}
}

func TestEncodeMap(t *testing.T) {
	tests := []struct {
		testMap  any
		expected string
	}{
		{
			map[string]string{"key1": "test"},
			`{ "key1": "test" }`,
		},
		{
			map[struct{ field int }]string{{field: 11}: "test3"},
			`{ { field: 11 }: "test3" }`,
		},
	}

	for _, tt := range tests {
		e := newEncoder()
		e.encodeMap(reflect.ValueOf(tt.testMap))
		res := e.String()
		if res != tt.expected {
			t.Errorf("encodeMap(%v) = %q expected %q", tt.testMap, res, tt.expected)
		}
	}
}

func TestEncoder(t *testing.T) {
	var sequel *string = nil

	movie := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actors: []Actor{
			{"Grp. Capt. Lionel Mandrake", "Peter Sellers"},
			{"Pres. Merkin Muffley", "Peter Sellers"},
			{"Gen. Buck Turgidson", "George C. Scott"},
			{"Brig. Gen. Jack D. Ripper", "Sterling Hayden"},
			{"Maj. T.J. \"King\" Kong", "Slim Pickens"},
			{"Dr. Strangelove", "Peter Sellers"},
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Sequel: sequel,
	}

	expected := `{ Title: "Dr. Strangelove", Subtitle: "How I Learned to Stop Worrying and Love the Bomb", Year: 1964, Actors: { { Role: "Grp. Capt. Lionel Mandrake", Name: "Peter Sellers" }, { Role: "Pres. Merkin Muffley", Name: "Peter Sellers" }, { Role: "Gen. Buck Turgidson", Name: "George C. Scott" }, { Role: "Brig. Gen. Jack D. Ripper", Name: "Sterling Hayden" }, { Role: "Maj. T.J. \"King\" Kong", Name: "Slim Pickens" }, { Role: "Dr. Strangelove", Name: "Peter Sellers" } }, Oscars: { "Best Actor (Nomin.)", "Best Adapted Screenplay (Nomin.)", "Best Director (Nomin.)", "Best Picture (Nomin.)" }, Sequel: nil }`
	res := Marshal(movie)

	if res != expected {
		t.Errorf("\nMarshal(%v)\nResponse: %q\nExpected: %q", movie, res, expected)
	}
}
