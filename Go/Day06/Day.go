//go:build d06

package Day06

import (
	AH "AoC2025/adventhelper"
	"strconv"
	"strings"
)

func parseInput(ss []string) ([][]int, string) {
	grid := [][]int{}
	ops := ""

	for i, s := range ss {
		if i != len(ss)-1 {
			row := []int{}
			ps := strings.Fields(s)
			for _, p := range ps {
				n, _ := strconv.Atoi(p)
				row = append(row, n)
			}
			grid = append(grid, row)
		} else {
			ps := strings.Fields(s)
			for _, r := range ps {
				ops += r
			}
		}
	}

	return grid, ops
}

func doCephalopodMaths(ss []string) int {
	rows, cols := len(ss), len(ss[0])
	op := ' '
	newSum := true
	bigSum := 0
	colValue := 0

	for col := 0; col < cols; col++ {
		if newSum {
			colValue = 0
			op = rune(ss[rows-1][col])
			if op == '*' {
				colValue = 1
			}
			newSum = false
		}
		number := 0
		for row := 0; row < rows-1; row++ {
			digit := 0
			rn := string(ss[row][col])
			if rn != " " {
				number *= 10
				digit, _ = strconv.Atoi(string(ss[row][col]))
			}
			number += digit
		}

		if number == 0 {
			bigSum += colValue
			newSum = true
		} else {
			if op == '+' {
				colValue += number
			} else {
				colValue *= number
			}
		}

		if col == cols-1 {
			bigSum += colValue
			newSum = true
		}
	}

	return bigSum
}

func doMaths(grid [][]int, ops string) int {
	total := 0
	rows := len(grid)

	for col, op := range ops {
		colt := 0
		if op == '*' {
			colt = 1
		}
		for row := 0; row < rows; row++ {
			if op == '+' {
				colt += grid[row][col]
			} else {
				colt *= grid[row][col]
			}
		}
		total += colt
	}

	return total
}

func Run() {
	is, _ := AH.ReadStrFile("../inputs/day06.txt")
	grid, ops := parseInput(is)

	p1 := doMaths(grid, ops)
	p2 := doCephalopodMaths(is)

	AH.PrintSoln(6, p1, p2)

	return
}
