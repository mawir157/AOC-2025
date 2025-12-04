//go:build d04
// +build d04

package Day04

import (
	AH "AoC2025/adventhelper"
)

type Pos struct {
	r, c int
}

func parseInput(ss []string) map[Pos]int {
	m := make(map[Pos]int)

	for r, s := range ss {
		for c, rn := range s {
			if rn == '@' {
				m[Pos{r, c}] = 0
			}
		}
	}

	return m
}

func nbrs(p Pos) [8]Pos {
	ns := [8]Pos{
		Pos{p.r - 1, p.c - 1},
		Pos{p.r - 1, p.c},
		Pos{p.r - 1, p.c + 1},
		Pos{p.r, p.c - 1},
		Pos{p.r, p.c + 1},
		Pos{p.r + 1, p.c - 1},
		Pos{p.r + 1, p.c},
		Pos{p.r + 1, p.c + 1},
	}

	return ns
}

func countNbrs(m map[Pos]int) {
	for p := range m {
		ns := nbrs(p)
		ni := 0
		for _, n := range ns {
			if _, ok := m[n]; ok {
				ni += 1
			}
		}
		m[p] = ni
	}
}
func clearWarehouse(m map[Pos]int) int {
	cleared := 0

	for true {
		xx := []Pos{}
		for p, n := range m {
			if n < 4 {
				xx = append(xx, p)
			}
		}
		cleared += len(xx)
		if len(xx) == 0 {
			break
		} else {
			for _, p := range xx {
				delete(m, p)
			}
			countNbrs(m)
		}
	}

	return cleared
}

func Run() {
	is, _ := AH.ReadStrFile("../inputs/day04.txt")
	wh := parseInput(is)
	countNbrs(wh)
	p1 := 0
	for _, n := range wh {
		if n < 4 {
			p1++
		}
	}
	p2 := clearWarehouse(wh)

	AH.PrintSoln(4, p1, p2)

	return
}
