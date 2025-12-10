//go:build d10

package Day10

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	AH "AoC2025/adventhelper"
)

func parseInput(s string) (int, []int, []int) {
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
	sort.Slice(moves, func(i, j int) bool {
		return AH.CountBits(moves[i]) > AH.CountBits(moves[j])
	})

	joltage := []int{}
	t = ss[len(ss)-1][1 : len(ss[len(ss)-1])-1]
	ns := strings.Split(t, ",")
	for _, n := range ns {
		v, _ := strconv.Atoi(n)
		joltage = append(joltage, v)
	}

	return target, moves, joltage
}

func groupTheoryInnit(t int, ms []int) int {
	// moves commute and are involutions, so the maximum possible sequence of
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

func applyMove(js []int, m int) []int {
	js_copy := make([]int, len(js))
	for i, v := range js {
		js_copy[i] = v
	}
	for i := 0; m > 0; i++ {
		if m&1 == 1 {
			js_copy[i] = js[i] - 1
		}
		m /= 2
	}

	return js_copy
}

var bestFound = 100000

func wtf(target []int, ms []int, presses int, mIdx int, st string) int {
	// again the moves commute and increment by one - so the move sequence can
	// be written as e.g. 00111222..nnn. Apply move 0 some number of times, then
	// move 1, moves 2 and so on.
	// Since the moves all increment the max number of moves is sum(joltage)
	// <hash|moves>
	if presses > 1000 {
		return 1000000
		panic("wtf")
	}

	// fmt.Println(presses, mIdx, target, st)

	if presses > bestFound || mIdx >= len(ms) {
		hi, lo := AH.MaxAndMin(target)
		if lo == 0 && hi == 0 {
			fmt.Println("FOUND A MATCH AT LENGTH", presses, st)
			if presses < bestFound {
				bestFound = presses
			}
			return bestFound
		}
		// fmt.Println("========================================")
		return 1000000
	}

	bestCount := 1000000
	newTarget := target
	m := ms[mIdx]

	// how many times can we apply m?
	iMax := 0
	for iMax = 0; ; iMax++ {
		newTarget = applyMove(newTarget, m)
		bad := false
		for _, v := range newTarget {
			if v < 0 {
				bad = true
			}
		}
		if bad {
			break
		}
	}
	// fmt.Println("Can apply", m, iMax, "times")
	newTarget = target

	for im := 0; im <= iMax; im++ {
		// fmt.Println("applying move", m, im, "times")
		// fmt.Println(target, "->", newTarget)
		k := wtf(newTarget, ms, presses+im, mIdx+1, st)

		if k < bestCount {
			bestCount = k
		}

		st += strconv.Itoa(m)
		st += ","

		newTarget = applyMove(newTarget, m)

		mxm, _ := AH.MaxAndMinIdx(newTarget)
		if presses+mxm+im >= bestFound {
			continue
		}
	}

	// for im, m := range ms {
	// 	if im < mIdx {
	// 		continue
	// 	}
	// 	newTarget := applyMove(target, m)
	// 	bad := false
	// 	maxTarget := 0
	// 	for _, v := range newTarget {
	// 		if v < 0 {
	// 			bad = true
	// 		}
	// 		if v > maxTarget {
	// 			maxTarget = v
	// 		}
	// 	}
	// 	if bad || presses+maxTarget >= bestFound {
	// 		continue
	// 	} else {
	// 		k := wtf(newTarget, ms, presses+1, im, memo)

	// 		if k < bestCount {
	// 			bestCount = k
	// 		}
	// 	}
	// }

	// fmt.Println("returning", bestFound)
	return bestCount
}

func Run() {
	// defer AH.TrackTime(time.Now(), "Day 10")
	is, _ := AH.ReadStrFile("../inputs/day10.txt")
	p1, p2 := 0, 0
	for i, s := range is {
		if i != 0 {
			continue
		}
		target, moves, joltage := parseInput(s)
		p1 += groupTheoryInnit(target, moves)
		fmt.Println("(", i, "/", len(is), ")", joltage)
		bestFound = 100000
		p2 += wtf(joltage, moves, 0, 0, "")
		fmt.Println(moves)
		// break
	}
	AH.PrintSoln(10, p1, p2)

	return
}
