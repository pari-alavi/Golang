package main

import (
	"reflect"
	"testing"
)

//1)Test the voting mechanism: Vote method
//2)Test the election results: GetResults method

func TestVote(t *testing.T) {

	election := Election{
		Ballots: []Ballot{
			{},
			{},
		},
	}

	//results
	// Vote in the first ballot
	election.Vote(0, "Candidate A")
	election.Vote(0, "Candidate B")
	election.Vote(0, "Candidate B")
	election.Vote(0, "Candidate C")
	election.Vote(0, "Candidate C")
	election.Vote(0, "Candidate C")
	election.Vote(0, "Candidate C")

	// Vote in the second ballot
	election.Vote(1, "Candidate A")
	election.Vote(1, "Candidate A")
	election.Vote(1, "Candidate B")
	election.Vote(1, "Candidate C")
	election.Vote(1, "Candidate C")
	election.Vote(1, "Candidate A")

	// Expected results after voting
	expected := Election{
		Ballots: []Ballot{
			{
				Candidates: []Candidate{
					{Name: "Candidate A", Votes: 1},
					{Name: "Candidate B", Votes: 2},
					{Name: "Candidate C", Votes: 4},
				},
			},
			{
				Candidates: []Candidate{
					{Name: "Candidate A", Votes: 3},
					{Name: "Candidate B", Votes: 1},
					{Name: "Candidate C", Votes: 2},
				},
			},
		},
	}

	// Check if the actual election matches the expected results

	if !reflect.DeepEqual(election, expected) {
		t.Errorf("Vote function did not produce the expected election results.\nExpected: %+v\nActual: %+v", expected, election)
	}

}

func TestGetResults(t *testing.T) {

	election := Election{
		Ballots: []Ballot{
			{},
			{},
		},
	}

	// Vote in the first ballot
	election.Vote(0, "Candidate A")
	election.Vote(0, "Candidate B")
	election.Vote(0, "Candidate B")
	election.Vote(0, "Candidate C")
	election.Vote(0, "Candidate C")
	election.Vote(0, "Candidate C")
	election.Vote(0, "Candidate C")

	// Vote in the second ballot
	election.Vote(1, "Candidate A")
	election.Vote(1, "Candidate A")
	election.Vote(1, "Candidate B")
	election.Vote(1, "Candidate C")
	election.Vote(1, "Candidate C")
	election.Vote(1, "Candidate A")

	// Expected result after voting
	expectedResults := []Candidate{
		{Name: "Candidate C", Votes: 6},
		{Name: "Candidate A", Votes: 4},
		{Name: "Candidate B", Votes: 3},
	}

	// Get the actual result using the GetResults method
	actualResults := election.GetResults()

	// Compare expected and actual results
	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Errorf("Expected %v, but got %v", expectedResults, actualResults)
	}
}
