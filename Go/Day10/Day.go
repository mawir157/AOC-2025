//go:build d10

package Day10

import (
	"strconv"
	"strings"
	"time"

	AH "AoC2025/adventhelper"
)

func parseInput(s string) (int, []int) {
	ss := strings.Split(s, " ")
	target := 0
	t := AH.ReverseString(ss[0][1 : len(ss[0])-1])
	for _, c := range t {
		target *= 2
		if c == '#' {
			target += 1
		}
	}

	moves := []int{}
	for i, s := range ss {
		if i == 0 || i == len(ss)-1 {
			continue
		}
		move := 0
		t := s[1 : len(s)-1]
		ns := strings.Split(t, ",")
		for _, n := range ns {
			v, _ := strconv.Atoi(n)
			move += (1 << v)
		}
		moves = append(moves, move)
	}

	return target, moves
}

func sortMoves(ms [][]int) [][]int {
	done := make([]bool, len(ms))
	toDo := len(ms)

	lengths := []int{}
	for _, m := range ms {
		l := 0
		for _, v := range m {
			l += v
		}
		lengths = append(lengths, l)
	}

	reorderedMoves := [][]int{}

	for toDo > 0 {
		target := make([]int, len(ms[0]))
		// find the joltage that appears in the fewest moves
		for i, m := range ms {
			if !done[i] {
				target = applyMove(target, m)
			}
		}
		// the joltage that fewest buttons affect and not already used
		_, aa := AH.MaxAndMin(target)
		mm := 0
		for idx, v := range target {
			if v == 0 {
				continue
			}
			if v >= aa {
				aa = v
				mm = idx
			}
		}

		// find a long move that affects button mm
		long := -1
		best := -1
		for i, m := range ms {
			if m[mm] != 0 && !done[i] {
				l := lengths[i]
				if l > long {
					long = l
					best = i
				}
			}
		}

		reorderedMoves = append(reorderedMoves, ms[best])
		done[best] = true
		toDo--

		if toDo == 1 {
			for i, b := range done {
				if !b {
					reorderedMoves = append(reorderedMoves, ms[i])
					break
				}
			}
			break
		}
	}

	return reorderedMoves
}

func moveSize(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func parseInput2(s string) ([]int, [][]int) {
	ss := strings.Split(s, " ")
	n := len(ss[0]) - 2

	moves := [][]int{}
	for i, s := range ss {
		if i == 0 || i == len(ss)-1 {
			continue
		}
		move := make([]int, n)
		t := s[1 : len(s)-1]
		ns := strings.Split(t, ",")
		for _, n := range ns {
			v, _ := strconv.Atoi(n)
			move[v] = 1
		}
		moves = append(moves, move)
	}

	moves = sortMoves(moves)

	joltage := []int{}
	t := ss[len(ss)-1][1 : len(ss[len(ss)-1])-1]
	ns := strings.Split(t, ",")
	for _, n := range ns {
		v, _ := strconv.Atoi(n)
		joltage = append(joltage, v)
	}

	return joltage, moves
}

func groupTheoryInnit(t int, ms []int) int {
	// moves commute are involutions, so the maximum possible sequence of
	// is len(ms) and every subsequence can be encoded w/ a binary number
	bestCountBits := 1000
	for bin := 0; bin < (1 << len(ms)); bin++ {
		s := 0
		bbin := bin
		countBits := 0
		for idx := 0; bbin > 0; idx++ {
			if bbin&1 == 1 {
				s ^= ms[idx]
				countBits++
			}
			bbin /= 2
		}

		if s == t && countBits < bestCountBits {
			bestCountBits = countBits
		}
	}
	return bestCountBits
}

func applyMove(js []int, m []int) []int {
	js_copy := make([]int, len(js))
	for i, v := range js {
		js_copy[i] = v - m[i]
	}

	return js_copy
}

var bestFound = 0

func wtf(target []int, moves [][]int, presses int, mIdx int) int {
	hi, _ := AH.MaxAndMin(target)

	if hi == 0 {
		bestFound = presses
		return bestFound
	}

	if presses > bestFound || mIdx >= len(moves) {
		return 1000000
	}

	bestCount := 1000000

	// unhittable jolts - we only have a subset of moves remaining.
	// If there are non-zero jolts that can no longer be hit this run is doomed
	unhit := make([]int, len(target))
	smallestMove := 100
	for idx := mIdx; idx < len(moves); idx++ {
		unhit = applyMove(unhit, moves[idx])
		smallestMove = AH.Min(smallestMove, len(moves[idx]))
	}

	for i, v := range unhit {
		if v == 0 && target[i] != 0 {
			return 1000000
		}
	}

	// remaining jolts
	totalJolt := 0
	largestJolt := 0
	for _, v := range target {
		totalJolt += v
		largestJolt = AH.Max(largestJolt, v)
	}
	if (bestFound-presses)*smallestMove < totalJolt {
		return 1000000
	}
	if (bestFound - presses) < largestJolt {
		return 1000000
	}

	for idx := mIdx; idx < len(moves); idx++ {
		m := moves[idx]
		newTarget := applyMove(target, m)
		_, lo := AH.MaxAndMin(newTarget)
		if lo < 0 || presses+1 >= bestFound {
			continue
		}
		k := wtf(newTarget, moves, presses+1, idx)

		if k < bestCount {
			bestCount = k
		}
	}

	return bestCount
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 10")
	is, _ := AH.ReadStrFile("../inputs/day10.txt")
	p1, p2 := 0, 0
	for i, s := range is {
		target, moves := parseInput(s)
		p1 += groupTheoryInnit(target, moves)
		joltage, moves2 := parseInput2(s)
		bestFound = 0
		for _, v := range joltage {
			bestFound += v
		}
		p2 += wtf(joltage, moves2, 0, 0)
	}
	AH.PrintSoln(10, p1, p2)

	return
}
