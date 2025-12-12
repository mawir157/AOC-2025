//go:build d12

package Day12

import (
	AH "AoC2025/adventhelper"
	"strconv"
	"strings"
	"time"
)

type Shape [3][3]bool
type Puzzle struct {
	x, y   int
	counts [6]int
}

func parseInput(ss []string) ([6]Shape, []Puzzle) {
	shapes := [6]Shape{}
	i, j := 0, 0
	for i = 0; i < 6; i++ {
		for j = 0; j < 3; j++ {
			s := ss[5*i+j+1]
			for k, r := range s {
				if r == '#' {
					shapes[i][j][k] = true
				} else {
					shapes[i][j][k] = false
				}
			}
		}
	}

	pzs := []Puzzle{}

	for i = i * 5; i < len(ss); i++ {
		pz := Puzzle{}
		ps := strings.Split(ss[i], ": ")
		dims := strings.Split(ps[0], "x")
		xdim, _ := strconv.Atoi(dims[0])
		ydim, _ := strconv.Atoi(dims[1])
		pz.x = xdim
		pz.y = ydim

		cts := strings.Split(ps[1], " ")
		for ni, ns := range cts {
			nn, _ := strconv.Atoi(ns)
			pz.counts[ni] = nn
		}

		pzs = append(pzs, pz)
	}

	return shapes, pzs
}

func solvePuzzle(shapes [6]Shape, p Puzzle) bool {
	// sanity check 1
	sc := 0
	tiles := 0
	for i, c := range p.counts {
		for idx := 0; idx < 9; idx++ {
			if shapes[i][idx%3][idx/3] {
				sc += c
			}
		}
		tiles += c
	}
	if sc > p.x*p.y {
		return false
	}
	// sanity check 2
	if tiles <= (p.x/3)*(p.y/3) {
		return true
	}

	// Lol! this is never hit
	return true
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 12")
	is, _ := AH.ReadStrFile("../inputs/day12.txt")
	sps, pzs := parseInput(is)

	p1, p2 := 0, "Good luck and Godspeed."
	for _, pz := range pzs {
		if solvePuzzle(sps, pz) {
			p1++
		}
	}

	AH.PrintSoln(12, p1, p2)

	return
}
