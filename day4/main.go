package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	defer f.Close()

	count := 0
	for scanner.Scan() {
		elf_ranges := strings.Split(scanner.Text(), ",")

		//Split ranges
		range_one := strings.Split(elf_ranges[0], "-")
		range_two := strings.Split(elf_ranges[1], "-")

		startOne, _ := strconv.Atoi(range_one[0])
		endOne, _ := strconv.Atoi(range_one[1])
		startTwo, _ := strconv.Atoi(range_two[0])
		endTwo, _ := strconv.Atoi(range_two[1])
		if(isOverlap(startOne,endOne,startTwo,endTwo)) {
			count += 1
		}


	}
	fmt.Println(count)
}
func isOverlap(range1Start, range1End, range2Start, range2End int) bool {
    // Check if either range is completely to the left or right of the other range
    if range1End < range2Start || range2End < range1Start {
        return false
    }
    
    // Otherwise, the ranges overlap
    return true
}

