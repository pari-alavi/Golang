package main

import (
	"fmt"
	"sort"
)

type Candidate struct {
	Name  string
	Votes int
}

func countAndSortVotes(ballots [][]string) []Candidate {
	// Initialize map to store candidate scores
	candidateScores := make(map[string]int)

	// Iterate over each ballot
	for _, ballot := range ballots {
		// Count votes for each candidate
		votes := make(map[string]int)
		for _, candidate := range ballot {
			votes[candidate]++
		}

		// Convert votes map to slice of candidates
		var candidates []Candidate
		for name, count := range votes {
			candidates = append(candidates, Candidate{Name: name, Votes: count})
		}

		// Sort candidates by votes using sort.Slice
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].Votes > candidates[j].Votes
		})

		// Assign scores 3, 2, 1 to top three candidates
		for i, candidate := range candidates {
			if i < 3 {
				candidateScores[candidate.Name] += (3 - i)
			}
		}
	}

	// Convert candidate scores map to slice of candidates
	var finalCandidates []Candidate
	for name, score := range candidateScores {
		finalCandidates = append(finalCandidates, Candidate{Name: name, Votes: score})
	}

	// Sort final candidates by scores using sort.Slice
	sort.Slice(finalCandidates, func(i, j int) bool {
		return finalCandidates[i].Votes > finalCandidates[j].Votes
	})

	return finalCandidates
}

func main() {
	// Sample ballots
	ballots := [][]string{
		{"A", "A", "A", "A", "B", "B", "B", "C", "C"},
		{"A", "A", "A", "A", "C", "C", "C", "B", "B"},
		{"C", "C", "C", "C", "A", "A", "A", "B", "B"},
	}

	// Call the function to count and sort votes
	finalCandidates := countAndSortVotes(ballots)

	// Display top 3 candidates with scores
	fmt.Println("Top 3 Candidates:")
	for i, candidate := range finalCandidates {
		if i < 3 {
			fmt.Printf("%s: %d Scores\n", candidate.Name, candidate.Votes)
		}
	}
}
