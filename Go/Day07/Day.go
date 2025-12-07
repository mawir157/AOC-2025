//go:build d07

package Day07

import (
	AH "AoC2025/adventhelper"
)

func parseInput(ss []string) (int, [][]int) {
	ts := [][]int{}
	startIdx := -1
	for r, s := range ss {
		if r%2 == 1 {
			continue
		}
		row := []int{}
		for idx, c := range s {
			if c == '^' {
				row = append(row, idx)
			}
			if c == 'S' {
				startIdx = idx
			}
		}
		if r > 0 {
			ts = append(ts, row)
		}
	}

	return startIdx, ts
}

func beamSplitter(start int, ss [][]int) (int, int) {
	splits := 0
	beam := make(map[int]int)
	beam[start] = 1

	for _, row := range ss {
		postSplit := make(map[int]int)

		for b := range beam {
			hasSplit := false
			for _, splitter := range row {
				if b == splitter {
					splits++
					postSplit[b-1] += beam[b]
					postSplit[b+1] += beam[b]
					hasSplit = true
					break
				}
			}
			if !hasSplit {
				postSplit[b] += beam[b]
			}
		}

		beam = postSplit
	}

	timeLines := 0
	for _, v := range beam {
		timeLines += v
	}

	return splits, timeLines
}

func Run() {
	is, _ := AH.ReadStrFile("../inputs/day07.txt")
	start, rs := parseInput(is)

	p1, p2 := beamSplitter(start, rs)

	AH.PrintSoln(7, p1, p2)

	return
}
