//go:build d01
// +build d01

package Day01

import (
	"strconv"

	AH "AoC2025/adventhelper"
)

func parseInput(ss []string) []int {
	moves := []int{}
	for _, s := range ss {
		c, n := s[0], s[1:]
		nn, _ := strconv.Atoi(n)

		if c == 'R' {
			moves = append(moves, nn)
		} else {
			moves = append(moves, -1*nn)

		}

	}

	return moves
}

func combination(ms []int) (int, int) {
	score1, score2 := 0, 0
	pos := 50
	for _, m := range ms {
		step := 1
		if m < 0 {
			step = -1
		}

		for c := 0; c != m; c += step {
			pos += step
			if pos%100 == 0 {
				score2++
			}
		}
		if pos%100 == 0 {
			score1++
		}
	}
	return score1, score2
}

func Run() {
	inputLines, _ := AH.ReadStrFile("../inputs/day01.txt")
	moves := parseInput(inputLines)
	p1, p2 := combination(moves)

	AH.PrintSoln(1, p1, p2)

	return
}
