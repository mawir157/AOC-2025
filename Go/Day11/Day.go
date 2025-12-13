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
	l, r int
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

func routes(adj Adj, ls Labels, from, to int) int {

	if v, ok := routeMemo[Pair{from, to}]; ok {
		return v
	}

	if from == to {
		return 1
	}

	paths := 0

	// find children of 'from'
	// fromIdx := ls[from]
	children := []int{}
	for i := 0; i < len(ls); i++ {
		if adj[from][i] == 1 {
			children = append(children, i)
		}
	}

	for _, p := range children {
		paths += routes(adj, ls, p, to)
	}

	routeMemo[Pair{from, to}] = paths
	return paths
}

func part2(adj Adj, ls Labels, from, to int) int {
	rts := 1
	ffti := ls["fft"]
	daci := ls["dac"]
	fft2dac := routes(adj, ls, ffti, daci)
	dac2fft := routes(adj, ls, daci, ffti)

	// one of these must be zero
	if fft2dac*dac2fft != 0 {
		panic("cyclic graph")
	}

	if fft2dac != 0 { // from -> fft -> dac -> to
		rts *= routes(adj, ls, from, ffti)
		rts *= fft2dac
		rts *= routes(adj, ls, daci, to)
	} else { // from -> dac -> fft -> to
		rts *= routes(adj, ls, from, daci)
		rts *= dac2fft
		rts *= routes(adj, ls, ffti, to)
	}

	return rts
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 11")
	is, _ := AH.ReadStrFile("../inputs/day11.txt")
	adj, lbls := buildGraph(is)
	p1 := routes(adj, lbls, lbls["you"], lbls["out"])
	p2 := part2(adj, lbls, lbls["svr"], lbls["out"])

	AH.PrintSoln(11, p1, p2)

	return
}
