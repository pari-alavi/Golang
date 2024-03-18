package main

import (
	"reflect"
	"testing"
)

func TestReversing(t *testing.T) {
	tests := []struct {
		input          string
		expectedResult string
	}{
		{
			input:          "This is a sample string with words of different lengths number 85962158",
			expectedResult: "This is a elpmas gnirts with sdrow of tnereffid shtgnel rebmun 85126958",
		},
		{
			input:          "",
			expectedResult: "",
		},
		{
			input:          "Hello",
			expectedResult: "olleH",
		},
		{
			input:          "12345",
			expectedResult: "54321",
		},
		{
			input:          "$#%!&",
			expectedResult: "&!%#$",
		},
	}

	for _, test := range tests {
		result := Reversing(test.input)
		if !reflect.DeepEqual(result, test.expectedResult) {
			t.Errorf("For input %q, expected %q, but got %q", test.input, test.expectedResult, result)
		}
	}
}
