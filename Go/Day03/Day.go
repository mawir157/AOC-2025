//go:build d03
// +build d03

package Day03

import (
	AH "AoC2025/adventhelper"
	"strconv"
)

func largestJoltage(s string, l int) int {
	to_drop := len(s) - l
	is := []int{}
	for _, r := range s {
		i, _ := strconv.Atoi(string(r))
		is = append(is, i)
	}
	jolt := 0

	for i := 0; i+to_drop < len(is); i++ {
		max, mIdx := AH.MaxAndMaxIdx(is[i : i+to_drop+1])

		jolt = 10*jolt + max

		i += mIdx
		to_drop -= mIdx
	}

	return jolt
}

func Run() {
	is, _ := AH.ReadStrFile("../inputs/day03.txt")
	p1, p2 := 0, 0
	for _, p := range is {
		p1 += largestJoltage(p, 2)
		p2 += largestJoltage(p, 12)
	}

	AH.PrintSoln(3, p1, p2)

	return
}
