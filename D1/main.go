package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var elfWithMostCalories []int

	in, err := os.Open("input.txt")
	defer in.Close()
	check(err)

	sc := bufio.NewScanner(in)

	currentElf := 0
	for sc.Scan() {
		line := sc.Text()
		
		if line == "" {
			// if currentElf > elfWithMostCalories {
			// 	elfWithMostCalories = currentElf
			// }
			elfWithMostCalories = append(elfWithMostCalories, currentElf)
			currentElf = 0
			continue
		}
		num, err := strconv.Atoi(line)
		check(err)
		currentElf += num
	}
	sort.Slice(elfWithMostCalories, func(i, j int) bool {
		return elfWithMostCalories[i] > elfWithMostCalories[j]
	})
	fmt.Println(elfWithMostCalories[0] + elfWithMostCalories[1] +elfWithMostCalories[2])

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}