package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func compareItems(inItemList [][]string) string {
	var result string
	compOne := inItemList[0]
	compTwo:= inItemList[1]
	compThree := inItemList[2]

	for _, two := range compTwo {
		for _, one := range compOne {
			if one == two {
				for _, three := range compThree {
					if three == one {
						result = one
					}
				}
			}
		}
	}
	return result
}

func getPriority(item string) int {
	var result int
	c := item[0]
	if unicode.IsUpper(rune(c)) {
		result = int(c)-38
	} else {
		result = int(c)-96
	}
	return result
}

func main() {
	var result int
	in, err := os.Open("input.txt")
	defer in.Close()
	check(err)
	var group [][]string
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		items := strings.Split(line, "")
		group = append(group, items)
		if len(group) == 3 {
			item := compareItems(group)
			if item != "" {
				result += getPriority(item)
			}
			group = nil
		}
	}
	fmt.Println("Result", result)
}