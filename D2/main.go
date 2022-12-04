package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkWin(a int, b int) bool {
	// A - Rock B - Paper C - Scissor
	// X - Rock Y - Paper Z - Scissor
	// A > Z
	// C > B
	// B > X
	if (a == 1 && b == 3) || (a == 3 && b == 2) || (a == 2 && b == 1) {
		return true
	}
	return false
}



func main() {
	// A - Rock B - Paper C - Scissor
	// X - Rock Y - Paper Z - Scissor
	// A > Z
	// C > B
	// B > X


	var totalPlayerOnePartOne, totalPlayerTwoPartOne int
	var totalPlayerOnePartTwo, totalPlayerTwoPartTwo int

	in, err := os.Open("input.txt")
	defer in.Close()
	check(err)

	sc := bufio.NewScanner(in)

	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		playerChose := strings.Split(line, " ")
		round1PlayerOne, round1PlayerTwo := scorePartOne(playerChose)
		totalPlayerOnePartOne += round1PlayerOne
		totalPlayerTwoPartOne += round1PlayerTwo

		round2PlayerOne, round2PlayerTwo := scorePartTwo(playerChose)
		totalPlayerOnePartTwo += round2PlayerOne
		totalPlayerTwoPartTwo += round2PlayerTwo

	}

	fmt.Println("PartOne", totalPlayerOnePartOne, totalPlayerTwoPartOne)
	fmt.Println("PartTwo", totalPlayerOnePartTwo, totalPlayerTwoPartTwo)
}

func scorePartOne(playerChose []string) (int, int){
	shapeScore := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	var totalPlayerOne int
	var totalPlayerTwo int
	if shapeScore[playerChose[0]] == shapeScore[playerChose[1]] {
		totalPlayerOne += 3
		totalPlayerTwo += 3
		fmt.Println("Draw", 6, 6)

	} else if checkWin(shapeScore[playerChose[0]], shapeScore[playerChose[1]]) {
		totalPlayerOne += 6
		fmt.Println("Player One Won", shapeScore[playerChose[0]]+6, shapeScore[playerChose[1]])
	} else if checkWin(shapeScore[playerChose[1]], shapeScore[playerChose[0]]) {
		totalPlayerTwo += 6
		fmt.Println("Player Two Won", shapeScore[playerChose[0]], shapeScore[playerChose[1]]+6)
	}
	totalPlayerOne += shapeScore[playerChose[0]]
	totalPlayerTwo += shapeScore[playerChose[1]]

	return totalPlayerOne, totalPlayerTwo
}


func scorePartTwo(playerChose []string) (int, int) {
	// A - Rock B - Paper C - Scissor
	// X - Rock Y - Paper Z - Scissor
	// A > Z - 3
	// C > B - 2
	// B > X - 1

	// A 3
	// B 1
	// C 2

	// A 2
	// B 3
	// C 1
	shapeScore := map[string]int{"A": 1, "B": 2, "C": 3}
	loseScore := map[string]int{"A": 3, "B": 1, "C": 2}
	winScore := map[string]int{"A": 2, "B": 3, "C": 1}

	var totalPlayerOne int
	var totalPlayerTwo int
	if playerChose[1] == "Y" {
		totalPlayerOne += 3 + shapeScore[playerChose[0]]
		totalPlayerTwo += 3 + shapeScore[playerChose[0]]
	} else if playerChose[1] == "X" {
		totalPlayerOne += 6 + shapeScore[playerChose[0]]
		totalPlayerTwo += loseScore[playerChose[0]]
	} else if playerChose[1] == "Z" {
		totalPlayerOne += shapeScore[playerChose[0]]
		totalPlayerTwo += 6 + winScore[playerChose[0]]
	}
	return totalPlayerOne, totalPlayerTwo
}
