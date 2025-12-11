//go:build d09

package Day09

import (
	AH "AoC2025/adventhelper"
	"strconv"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

type Edge struct {
	p, q Pos
	vert bool
}

func parseInput(ss []string) ([]Pos, []Edge) {
	ps := []Pos{}

	for _, p := range ss {
		pr := strings.Split(p, ",")
		x, _ := strconv.Atoi(pr[0])
		y, _ := strconv.Atoi(pr[1])
		ps = append(ps, Pos{x, y})
	}

	es := []Edge{}

	for i := range ps {
		if i == len(ps)-1 {
			es = append(es, Edge{ps[i], ps[0], ps[i].x == ps[0].x})
		} else {
			es = append(es, Edge{ps[i], ps[i+1], ps[i].x == ps[i+1].x})
		}
	}

	return ps, es
}

func maxRect(ps []Pos, es []Edge) (int, int) {
	maxRect1, maxRect2 := 0, 0
	for iq := 0; iq < len(ps); iq++ {
		q := ps[iq]
		for ip := iq + 1; ip < len(ps); ip++ {
			p := ps[ip]

			area := (AH.AbsInt(q.x-p.x) + 1) * (AH.AbsInt(q.y-p.y) + 1)

			if area > maxRect1 {
				maxRect1 = area
			}

			if area > maxRect2 {
				if goodRect(p, q, es) {
					maxRect2 = area
				}
			}
		}
	}

	return maxRect1, maxRect2
}

func goodRect(v1, v2 Pos, es []Edge) bool {
	if v1.x == v2.x || v1.y == v2.y {
		return false // this rectangle is line, so is unlikely to have max area
	}

	good := true

	for _, e := range es {
		if e.vert {
			xlo, xhi := AH.Min(v1.x, v2.x), AH.Max(v1.x, v2.x)

			if xlo < e.p.x && e.p.x < xhi {
				for y := e.p.y; y != e.q.y; y += AH.Sign(e.q.y - e.p.y) {
					if (y-v1.y)*(y-v2.y) < 0 {
						good = false
						break
					}
				}
			}
		} else {
			ylo, yhi := AH.Min(v1.y, v2.y), AH.Max(v1.y, v2.y)

			if ylo < e.p.y && e.p.y < yhi {
				for x := e.p.x; x != e.q.x; x += AH.Sign(e.q.x - e.p.x) {
					if (x-v1.x)*(x-v2.x) < 0 {
						good = false
						break
					}
				}
			}
		}

		if !good {
			break
		}
	}

	return good
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 9")
	is, _ := AH.ReadStrFile("../inputs/day09.txt")
	ps, es := parseInput(is)
	p1, p2 := maxRect(ps, es)

	AH.PrintSoln(9, p1, p2)

	return
}
