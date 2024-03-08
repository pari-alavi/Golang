package main

import (
	"fmt"
	"sort"
)

// Candidate represents a candidate in the ballot
type Candidate struct {
	Name  string
	Votes int
}

// Ballot struct represents the list of candidates
type Ballot struct {
	Candidates []Candidate
}

// Election struct represents multiple ballots
type Election struct {
	Ballots []Ballot
}

// Vote function increments the vote count for the selected candidate
func (e *Election) Vote(ballotIndex int, candidateName string) {
	if ballotIndex >= 0 && ballotIndex < len(e.Ballots) {
		e.Ballots[ballotIndex].Vote(candidateName)
	}
}
func (b *Ballot) Vote(candidateName string) {
	for i := range b.Candidates {
		if b.Candidates[i].Name == candidateName {
			b.Candidates[i].Votes++
			return
		}
	}
	b.Candidates = append(b.Candidates, Candidate{Name: candidateName, Votes: 1})
}

// GetResults function aggregates and returns the election results
func (e *Election) GetResults() []Candidate {
	// Aggregate votes across all ballots
	aggregateResults := make(map[string]int)
	for _, ballot := range e.Ballots {
		for _, candidate := range ballot.Candidates {
			aggregateResults[candidate.Name] += candidate.Votes
		}
	}

	// Convert the aggregated results to a slice of candidates
	var results []Candidate
	for name, votes := range aggregateResults {
		results = append(results, Candidate{Name: name, Votes: votes})
	}

	// Sort candidates based on votes in descending order.
	sort.Slice(results, func(i, j int) bool {
		return results[i].Votes > results[j].Votes
	})

	return results
}

func main() {
	// // Create an election with multiple ballots
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

	results := election.GetResults()
	fmt.Println("Election Results:")
	for _, candidate := range results {
		fmt.Printf("%s: %d votes\n", candidate.Name, candidate.Votes)
	}
}
