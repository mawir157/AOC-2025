//go:build d11

package Day11

import (
	AH "AoC2025/adventhelper"
	"strings"
	"time"
)

type Adj [][]int
type Labels map[string]int
type Pair struct {
	l, r string
}

var routeMemo = make(map[Pair]int)

func buildGraph(ss []string) (Adj, Labels) {
	label := 0
	map_labels := make(Labels)

	// get all the labels
	for _, s := range ss {
		ps := strings.Split(s, ": ")

		if _, ok := map_labels[ps[0]]; !ok {
			map_labels[ps[0]] = label
			label++
		}

		qs := strings.Split(ps[1], " ")
		for _, q := range qs {
			if _, ok := map_labels[q]; !ok {
				map_labels[q] = label
				label++
			}
		}

	}
	adj := make([][]int, len(map_labels))
	for i := range adj {
		adj[i] = make([]int, len(map_labels))
	}

	for _, s := range ss {
		ps := strings.Split(s, ": ")
		from := map_labels[ps[0]]

		qs := strings.Split(ps[1], " ")
		for _, q := range qs {
			to := map_labels[q]
			adj[from][to] = 1
		}
	}

	return adj, map_labels
}

func routesDown(adj Adj, ls Labels, from, to string) int {

	if v, ok := routeMemo[Pair{from, to}]; ok {
		return v
	}

	if from == to {
		return 1
	}

	paths := 0

	// find children of 'from'
	fromIdx := ls[from]
	children := []int{}
	for i := 0; i < len(ls); i++ {
		if adj[fromIdx][i] == 1 {
			children = append(children, i)
		}
	}

	for _, p := range children {
		childStr := ""
		for k, v := range ls {
			if v == p {
				childStr = k
			}
		}
		if len(childStr) == 0 {
			panic("CAN'T FIND PARENT")
		}

		paths += routesDown(adj, ls, childStr, to)
	}

	routeMemo[Pair{from, to}] = paths
	return paths
}

func part2(adj Adj, ls Labels, from, to string) int {
	rts := 1
	fft2dac := routesDown(adj, ls, "fft", "dac")
	dac2fft := routesDown(adj, ls, "dac", "fft")

	// one of these must be zero
	if fft2dac*dac2fft != 0 {
		panic("cyclic graph")
	}

	if fft2dac != 0 { // from -> fft -> dac -> to
		rts *= routesDown(adj, ls, from, "fft")
		rts *= fft2dac
		rts *= routesDown(adj, ls, "dac", to)
	} else { // from -> dac -> fft -> to
		rts *= routesDown(adj, ls, from, "dac")
		rts *= dac2fft
		rts *= routesDown(adj, ls, "fft", to)
	}

	return rts
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 11")
	is, _ := AH.ReadStrFile("../inputs/day11.txt")
	adj, lbls := buildGraph(is)
	p1 := routesDown(adj, lbls, "you", "out")
	p2 := part2(adj, lbls, "svr", "out")

	AH.PrintSoln(11, p1, p2)

	return
}
