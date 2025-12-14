//go:build d10

package Day10

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	AH "AoC2025/adventhelper"
)

type Move []int
type Joltage []int

type Hash struct {
	q1, q2, q3 int
}

var INFINITY = 1000000

var memo = make(map[Hash]int)

func moveSize(arr Move) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func hashJoltage(js Joltage) Hash {
	// assume joltage < 512
	q1, q2, q3 := 0, 0, 0
	for i, j := range js {
		if i < 4 {
			q1 *= 512
			q1 += j
		} else if i < 8 {
			q2 *= 512
			q2 += j
		} else if i < 12 {
			q3 *= 512
			q3 += j
		}
	}

	return Hash{q1, q2, q3}
}

func parseInput2(s string) (Joltage, Joltage, []Move) {
	ss := strings.Split(s, " ")
	n := len(ss[0]) - 2

	t := ss[0][1 : len(ss[0])-1]
	joltage1 := make(Joltage, len(t))
	for i, c := range t {
		if c == '#' {
			joltage1[i] = 1
		}
	}

	moves := []Move{}
	for i, s := range ss {
		if i == 0 || i == len(ss)-1 {
			continue
		}
		move := make(Move, n)
		t := s[1 : len(s)-1]
		ns := strings.Split(t, ",")
		for _, n := range ns {
			v, _ := strconv.Atoi(n)
			move[v] = 1
		}
		moves = append(moves, move)
	}

	// sort.Slice(moves, func(i, j int) bool {
	// 	return moveSize(moves[i]) > moveSize(moves[j])
	// })

	joltage2 := Joltage{}
	t = ss[len(ss)-1][1 : len(ss[len(ss)-1])-1]
	ns := strings.Split(t, ",")
	for _, n := range ns {
		v, _ := strconv.Atoi(n)
		joltage2 = append(joltage2, v)
	}

	return joltage1, joltage2, moves
}

func applyMove(js Joltage, move Move, part2 bool) Joltage {
	js_copy := make(Joltage, len(js))
	for i, v := range js {
		if part2 {
			js_copy[i] = v - move[i]
		} else {
			js_copy[i] = (v + move[i]) % 2
		}
	}

	return js_copy
}

func allEven(arr Joltage) bool {
	for _, v := range arr {
		if v%2 == 1 {
			return false
		}
	}

	return true
}

func binMap(moves []Move) map[int]Move {
	bm := make(map[int]Move)
	for bin := 0; bin < (1 << len(moves)); bin++ {
		bbin := bin
		move := make([]int, len(moves[0]))
		for idx := 0; bbin > 0; idx++ {
			if bbin&1 == 1 {
				mv := moves[idx]
				move = applyMove(move, mv, true)
			}
			bbin /= 2
		}
		for i := range move {
			move[i] *= -1
		}
		bm[bin] = move
	}
	// purge redundant combos i.e two sequences of moves lead to the same
	tokill := []int{}
	for bin1, mv1 := range bm {
		for bin2, mv2 := range bm {
			if bin1 >= bin2 {
				continue
			}
			bad := true
			for i := range mv2 {
				if mv1[i] != mv2[i] {
					bad = false
					break
				}
			}
			if bad {
				if AH.CountBits(bin1) < AH.CountBits(bin2) {
					tokill = append(tokill, bin2)
				} else if AH.CountBits(bin2) < AH.CountBits(bin1) {
					tokill = append(tokill, bin1)
				}
			}
		}
	}
	for _, kill := range tokill {
		delete(bm, kill)
	}

	return bm
}

// part2 = true only care if resulting joltage is even
// false joltage must be zero
func groupTheory(joltage Joltage, mm map[int]Move, part2 bool) []int {
	validSequences := []int{}
	for bin, mv := range mm {
		jtg := applyMove(joltage, mv, part2)
		hi, lo := AH.MaxAndMin(jtg)
		if part2 {
			if allEven(jtg) && hi >= 0 && lo >= 0 {
				validSequences = append(validSequences, bin)
			}
		} else {
			if hi == 0 && lo == 0 {
				validSequences = append(validSequences, bin)
			}
		}
	}
	return validSequences
}

func part1(joltage Joltage, mm map[int]Move) int {
	val := INFINITY
	vs := groupTheory(joltage, mm, false)
	for _, v := range vs {
		vv := AH.CountBits(v)
		if val > vv {
			val = vv
		}
	}

	return val
}

func formatDebugArr(debugArr []int) {
	m := [24]int{}
	total := 0
	for _, n := range debugArr {
		for i := 0; n > 0; i++ {
			if n&1 == 1 {
				m[i]++
				total++
			}
			n /= 2
		}
	}
	fmt.Println(m, total)
	return
}

func wtf2(joltage Joltage, mm map[int]Move, first bool) int {
	hash := hashJoltage(joltage)
	if v, ok := memo[hash]; ok {
		return v
	}

	hi, lo := AH.MaxAndMin(joltage)

	if hi == 0 && lo == 0 {
		return 0
	}

	// only relevant if joltage is initially all even
	evenInit := first && allEven(joltage)

	possibleMoves := groupTheory(joltage, mm, true)
	if evenInit {
		possibleMoves = []int{}
		for k, _ := range mm {
			possibleMoves = append(possibleMoves, k)
		}
	}

	if len(possibleMoves) == 0 {
		memo[hash] = INFINITY
		return INFINITY
	}

	best := INFINITY

	for _, m := range possibleMoves {
		jtg := make([]int, len(joltage))
		for i := range joltage {
			jtg[i] = joltage[i]
		}
		jtg = applyMove(jtg, mm[m], true)
		count := AH.CountBits(m)
		if evenInit {
			jtg = applyMove(jtg, mm[m], true)
			count += AH.CountBits(m)
		}
		// at this point next Joltage is all even
		hi, lo := AH.MaxAndMin(jtg)
		if lo < 0 {
			continue
		}

		if !allEven(jtg) {
			panic("not recovering from here")
		}
		multiplier := 1
		for allEven(jtg) && hi > 0 {
			for i := range jtg {
				jtg[i] /= 2
			}
			multiplier *= 2
			break
		}

		t := count + multiplier*wtf2(jtg, mm, false)
		if best > t {
			best = t
		}
	}

	memo[hash] = best
	return best
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 10")
	is, _ := AH.ReadStrFile("../inputs/day10.txt")
	p1, p2 := 0, 0
	for _, s := range is {
		j1, j2, moves := parseInput2(s)
		bm := binMap(moves)
		p1 += part1(j1, bm)
		memo = make(map[Hash]int)
		p2 += wtf2(j2, bm, true)
	}
	AH.PrintSoln(10, p1, p2)

	return
}
