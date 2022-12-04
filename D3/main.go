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

func compareItems(inItemList []string) string {
	var result string
	length := len(inItemList) / 2
	compOne := inItemList[0:length]
	compTwo:= inItemList[length:]
	for _, two := range compTwo {
		for _, one := range compOne {
			if one == two {
				result = one
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

	sc := bufio.NewScanner(in)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		items := strings.Split(line, "")
		item := compareItems(items)
		result += getPriority(item)
	}
	fmt.Println("Result", result)
}