package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	defer f.Close()

	scanner := bufio.NewScanner(f)

	cals := []int{}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == "\n" {
			cals = append(cals, total)
			total = 0
		} else {
			num, _ := strconv.Atoi(line)
			total += num
		}
	}
	sort.Ints(cals)
	fmt.Println(cals[len(cals)-1] + cals[len(cals)-2] + cals[len(cals)-3])
}
