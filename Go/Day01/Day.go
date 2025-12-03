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

func combination(pos, vals int, ms []int) (int, int) {
	score1, score2 := 0, 0
	for _, m := range ms {
		s, e := pos, pos+m
		if e%vals == 0 {
			score1++
		}
		if s < e {
			e = AH.FloorDiv(e, vals)
			s = AH.FloorDiv(s, vals)
		} else {
			e = AH.FloorDiv(e-1, vals)
			s = AH.FloorDiv(s-1, vals)
		}
		score2 += AH.AbsInt(e - s)
		pos += m
	}
	return score1, score2
}

func Run() {
	inputLines, _ := AH.ReadStrFile("../inputs/day01.txt")
	moves := parseInput(inputLines)
	p1, p2 := combination(50, 100, moves)

	AH.PrintSoln(1, p1, p2)

	return
}
