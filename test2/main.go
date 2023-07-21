package main

import (
	"fmt"
	"sort"
)

func denseRanking(scores []int, newScores []int) []int {
	uniqueScores := make([]int, 0)
	scoreMap := make(map[int]bool)
	for _, score := range scores {
		if _, ok := scoreMap[score]; !ok {
			scoreMap[score] = true
			uniqueScores = append(uniqueScores, score)
		}
	}
	sort.Slice(uniqueScores, func(i, j int) bool {
		return uniqueScores[i] > uniqueScores[j]
	})

	rankings := make([]int, len(newScores))
	for i, newScore := range newScores {
		rank := 1
		for _, score := range uniqueScores {
			if newScore >= score {
				break
			}
			rank++
		}
		rankings[i] = rank
	}

	return rankings
}

func main() {
	// Sample input 1
	scores1 := []int{100, 100, 50, 40, 40, 20, 10}
	newScores1 := []int{5, 25, 50, 120}

	rankings1 := denseRanking(scores1, newScores1)

	fmt.Println("Sample input 1:")
	fmt.Println("Hasil peringkat baru untuk skor baru:")
	for i, rank := range rankings1 {
		fmt.Printf("Skor %d: Peringkat %d\n", newScores1[i], rank)
	}
	fmt.Println()

	// Sample input 2
	scores2 := []int{100, 80, 75, 70, 40, 40, 30}
	newScores2 := []int{100, 85, 40}

	rankings2 := denseRanking(scores2, newScores2)

	fmt.Println("Sample input 2:")
	fmt.Println("Hasil peringkat baru untuk skor baru:")
	for i, rank := range rankings2 {
		fmt.Printf("Skor %d: Peringkat %d\n", newScores2[i], rank)
	}
	fmt.Println()

	// Sample input 3
	scores3 := []int{100, 60, 40, 30, 20}
	newScores3 := []int{70, 50, 20, 10}

	rankings3 := denseRanking(scores3, newScores3)

	fmt.Println("Sample input 3:")
	fmt.Println("Hasil peringkat baru untuk skor baru:")
	for i, rank := range rankings3 {
		fmt.Printf("Skor %d: Peringkat %d\n", newScores3[i], rank)
	}
}
