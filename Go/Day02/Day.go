//go:build d02

package Day02

import (
	AH "AoC2025/adventhelper"
	"strconv"
	"strings"
	"time"
)

type Pair struct {
	lhs, rhs int
}

func parseInput(s string) []Pair {
	ps := []Pair{}
	ss := strings.Split(s, ",")

	for _, p := range ss {
		pr := strings.Split(p, "-")
		lhs, _ := strconv.Atoi(pr[0])
		rhs, _ := strconv.Atoi(pr[1])
		ps = append(ps, Pair{lhs, rhs})
	}

	return ps
}

func countDigits(i int) int {
	digits := 0
	for i > 0 {
		digits++
		i /= 10
	}

	return digits
}

func countInvalid(p Pair) (int, int) {
	count1, count2 := 0, 0
	for i := p.lhs; i <= p.rhs; i++ {
		d := countDigits(i)
		for reps := 2; reps <= d; reps++ {
			// prevent leading zeros!
			if d%reps != 0 {
				continue
			}
			ii := i
			mask := AH.PowInt(10, d/reps)
			is_rep_unit := true
			rep_block := ii % mask
			ii /= mask
			for ii > 0 {
				if ii%mask != rep_block {
					is_rep_unit = false
					break
				}
				ii /= mask
			}

			if is_rep_unit {
				if reps == 2 {
					count1 += i
				}
				count2 += i
				break
			}
		}
	}
	return count1, count2
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 2")
	is, _ := AH.ReadStrFile("../inputs/day02.txt")
	ps := parseInput(is[0])
	p1, p2 := 0, 0
	for _, p := range ps {
		v1, v2 := countInvalid(p)
		p1 += v1
		p2 += v2
	}

	AH.PrintSoln(2, p1, p2)

	return
}
