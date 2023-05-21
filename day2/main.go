package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Rock paper scissors <-- order
	other_options := []string{"A", "B", "C"}
	// Lose draw win
	my_options := []string{"X", "Y", "Z"}
	scores := [][]int{
		{3, 4, 8},
		{1, 5, 9},
		{2, 6, 7},
	}
	score := 0

	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	defer f.Close()

	for scanner.Scan() {

		line := strings.Split(scanner.Text(), " ")

		my_choice := line[1]
		other_choice := line[0]
		my_index := findIndex(my_choice, my_options)
		other_index := findIndex(other_choice, other_options)

		score += scores[other_index][my_index]

	}
	fmt.Println(score)
}
func findIndex(choice string, options []string) int {
	for i, option := range options {
		if option == choice {
			return i
		}
	}
	return -1
}
