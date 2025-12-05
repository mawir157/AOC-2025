//go:build d05
// +build d05

package Day05

import (
	AH "AoC2025/adventhelper"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	lo, hi int
}

func parseInput(ss []string) ([]Interval, []int) {
	rs := []Interval{}
	ns := []int{}
	intervals := true
	for _, s := range ss {
		if len(s) == 0 {
			intervals = false
		}
		if intervals {
			pr := strings.Split(s, "-")
			lhs, _ := strconv.Atoi(pr[0])
			rhs, _ := strconv.Atoi(pr[1])
			rs = append(rs, Interval{lhs, rhs})
		} else {
			nstr, _ := strconv.Atoi(s)
			ns = append(ns, nstr)
		}
	}

	return rs, ns
}

func checkIngredients(rs []Interval, ns []int) int {
	allGood := 0

	for _, n := range ns {
		good := false
		for _, i := range rs {
			if i.lo <= n && n <= i.hi {
				good = true
				break
			}
		}
		if good {
			allGood++
		}
	}

	return allGood
}

func checkIngredientsAlt(rs []Interval) int {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].lo < rs[j].lo
	})

	merged_rs := []Interval{}

	A := rs[0]
	for idx := 1; idx < len(rs); idx++ {
		B := rs[idx]
		if A.hi < B.lo { // no overlap insert A and move onto B
			merged_rs = append(merged_rs, A)
			A = B
		} else if A.hi > B.hi { // A completely contains B, so just keep A
			continue
		} else { // A and B over lap so merge
			A = Interval{A.lo, B.hi}
		}
	}
	// remember to add final interval
	merged_rs = append(merged_rs, A)

	total := 0
	for _, r := range merged_rs {
		total += r.hi - r.lo + 1
	}

	return total
}

func Run() {
	is, _ := AH.ReadStrFile("../inputs/day05.txt")
	rs, ns := parseInput(is)

	p1 := checkIngredients(rs, ns)
	p2 := checkIngredientsAlt(rs)

	AH.PrintSoln(5, p1, p2)

	return
}
