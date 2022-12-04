package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getRange(r string) []int {
	min, max := getMinMax(r)
	result := make([]int, max-min+1)
	for i := range result {
		result[i] = min + i
	}
	return result
}

func getMinMax(r string) (int, int) {
	rList := strings.Split(r, "-")
	min, err := strconv.Atoi(rList[0])
	check(err)
	max, err := strconv.Atoi(rList[1])
	check(err)
	fmt.Println(min, max)
	return min, max
}

func checkContains(min1, max1, min2, max2 int) bool {
	if (min2 >= min1 && max2 <= max1) || (min1 >= min2 && max1 <= max2) {
		return true
	}
	return false
}

func main() {
	in, err := os.Open("input.txt")
	check(err)
	var s int

	sc := bufio.NewScanner(in)
	for sc.Scan() {
		line := sc.Text()
		pairs := strings.Split(line, ",")
		min1, max1 := getMinMax(pairs[0])
		min2, max2 := getMinMax(pairs[1])
		if checkContains(min1, max1, min2, max2) {
			s += 1
		}
	}
	fmt.Println("Partone: ", s)
}