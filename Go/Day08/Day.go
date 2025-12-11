//go:build d08

package Day08

import (
	AH "AoC2025/adventhelper"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Pos struct {
	x, y, z int
}

func parseInput(ss []string) []Pos {
	ps := []Pos{}

	for _, s := range ss {
		pr := strings.Split(s, ",")
		x, _ := strconv.Atoi(pr[0])
		y, _ := strconv.Atoi(pr[1])
		z, _ := strconv.Atoi(pr[2])
		ps = append(ps, Pos{x, y, z})
	}

	return ps
}

func buildGraph(ps []Pos, cons int) [][]bool {
	adj := make([][]bool, len(ps))
	for i := range ps {
		adj[i] = make([]bool, len(ps))
	}

	pIdx, qIdx := -1, -1
	for i := 0; i < cons; i++ {
		dist := 1000000000000000
		for iq := 0; iq < len(ps); iq++ {
			q := ps[iq]
			for ip := iq + 1; ip < len(ps); ip++ {
				p := ps[ip]
				if adj[iq][ip] {
					continue
				}

				distSqrd := (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y) + (p.z-q.z)*(p.z-q.z)
				if distSqrd < dist {
					pIdx = ip
					qIdx = iq
					dist = distSqrd
				}
			}
		}
		adj[qIdx][pIdx] = true
		adj[pIdx][qIdx] = true
	}

	return adj
}

func findCCs(adj [][]bool) int {
	comps := [][]int{}
	flagged := make([]bool, len(adj))
	count := 0
	open := 0
	for i, b := range flagged {
		if !b {
			count++
			open = i
		}
	}

	for count > 0 {
		comp := []int{}
		count = 0

		queue := make(map[int]bool)
		queue[open] = true
		flagged[open] = true
		comp = append(comp, open)

		for len(queue) > 0 {
			p := 0
			for k, _ := range queue {
				p = k
				break
			}
			delete(queue, p)
			for q, b := range adj[p] {
				if b && !flagged[q] {
					comp = append(comp, q)
					flagged[q] = true
					queue[q] = true
				}
			}
		}

		for i, b := range flagged {
			if !b {
				count++
				open = i
			}
		}
		comps = append(comps, comp)

	}
	cs := []int{}
	for _, c := range comps {
		cs = append(cs, len(c))
	}
	sort.Slice(cs, func(i, j int) bool {
		return cs[i] > cs[j]
	})

	return cs[0] * cs[1] * cs[2]
}

func mostIsolatedVertex(ps []Pos) int {
	nns := make([]int, len(ps))
	nps := make([]int, len(ps))
	for ip, p := range ps {
		nn := 1000000000
		for iq, q := range ps {
			if iq == ip {
				continue
			}
			distSqrd := (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y) + (p.z-q.z)*(p.z-q.z)
			if distSqrd < nn {
				nn = distSqrd
				nps[ip] = iq
			}
		}
		nns[ip] = nn
	}

	_, i := AH.MaxAndMaxIdx(nns)

	return ps[i].x * ps[nps[i]].x
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 8")
	is, _ := AH.ReadStrFile("../inputs/day08.txt")
	ps := parseInput(is)
	adj := buildGraph(ps, 1000)
	p1 := findCCs(adj)

	p2 := mostIsolatedVertex(ps)

	AH.PrintSoln(8, p1, p2)

	return
}
