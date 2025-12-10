//go:build d10

package Day10

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	AH "AoC2025/adventhelper"
)

func parseInput(s string) ([]bool, [][]int) {
	ss := strings.Split(s, " ")
	target := []bool{}
	t := ss[0][1 : len(ss[0])-1]
	fmt.Println(t)
	for _, c := range t {
		fmt.Println(c)
		if c == '#' {
			target = append(target, true)
		} else if c == '.' {
			target = append(target, false)
		} else {
			panic("Indices are wrong")
		}
	}
	moves := [][]int{}
	for i, s := range ss {
		if i == 0 || i == len(ss)-1 {
			continue
		}
		move := []int{}
		t := s[1 : len(s)-1]
		ns := strings.Split(t, ",")
		for _, n := range ns {
			v, _ := strconv.Atoi(n)
			move = append(move, v)
		}
		moves = append(moves, move)
	}

	return target, moves
}

func equal(lhs []bool, rhs []bool) {
	for i, b := range lhs {
		if b != rhs[i] {
			return false
		}
	}
	return true
}

func playGame(target []bool, moves [][]int, state []bool) int {
	if equal(target, state) {
		return 1
	}
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 10")
	inputLines, _ := AH.ReadStrFile("../inputs/day10.txt")
	target, moves := parseInput(inputLines[0])
	fmt.Println(target)
	fmt.Println(moves)
	p1, p2 := 0, 0

	AH.PrintSoln(10, p1, p2)

	return
}
