package main

import (
	"reflect"
	"testing"
)

//unit test for the TestCountAndSortVotes Function

func TestCountAndSortVotes(t *testing.T) {

	//passing input to the function
	ballots := [][]string{
		{"A", "A", "A", "A", "B", "B", "B", "C", "C"},
		{"A", "A", "A", "A", "C", "C", "C", "B", "B"},
		{"C", "C", "C", "C", "A", "A", "A", "B", "B"},
	}
	// Get the actual result using the CountAndSortVotes function
	results := countAndSortVotes(ballots)

	// Expected results of the function
	expected := []Candidate{
		{Name: "A", Votes: 8},
		{Name: "C", Votes: 6},
		{Name: "B", Votes: 4},
	}

	// Compare expected and actual results
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, results)

	}
}
