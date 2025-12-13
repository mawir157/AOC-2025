//go:build d10

package Day10

import (
	"fmt"
	"sort"
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

	sort.Slice(moves, func(i, j int) bool {
		return moveSize(moves[i]) > moveSize(moves[j])
	})

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
	return bm
}

// part2 = true only care if resulting joltage is even
// false joltage must be zero
func groupTheory(joltage Joltage, mm map[int]Move, part2 bool) []int {
	validSequences := []int{}
	for bin, mv := range mm {
		jtg := applyMove(joltage, mv, part2)
		if part2 {
			if allEven(jtg) {
				validSequences = append(validSequences, bin)
			}
		} else {
			hi, lo := AH.MaxAndMin(jtg)
			if hi == 0 && lo == 0 {
				validSequences = append(validSequences, bin)
			}
		}
	}
	return validSequences
}

func part1(joltage Joltage, mm map[int]Move) int {
	val := 1000000
	vs := groupTheory(joltage, mm, false)
	for _, v := range vs {
		vv := AH.CountBits(v)
		if val > vv {
			val = vv
		}
	}

	return val
}

func wtf2(joltage Joltage, mm map[int]Move) int {
	// fmt.Println(joltage)
	collide, debug := false, 0
	hash := hashJoltage(joltage)
	if v, ok := memo[hash]; ok {
		collide, debug = true, v
		return v
	}

	hi, _ := AH.MaxAndMin(joltage)

	if hi == 0 {
		return 0
	}

	possibleMoves := groupTheory(joltage, mm, true)
	// fmt.Println(possibleMoves)
	if len(possibleMoves) == 0 {
		memo[hash] = 1000000
		return 1000000
	}

	best := 1000000
	for _, m := range possibleMoves {
		jtg := make([]int, len(joltage))
		for i := range joltage {
			jtg[i] = joltage[i]
		}
		// fmt.Println(jtg, joltage, m, mm[m])
		jtg = applyMove(jtg, mm[m], true)
		// fmt.Println(jtg)
		count := AH.CountBits(m)
		// at this point next Joltage is all even
		for i := range jtg {
			jtg[i] /= 2
		}
		// fmt.Println(count, jtg)

		_, lo := AH.MaxAndMin(jtg)
		if lo < 0 {
			continue
		}
		// break
		t := count + 2*wtf2(jtg, mm)
		if best > t {
			best = t
			// if presses == 0 {
			// 	fmt.Println("new best", best)
			// }
		}
	}

	if collide {
		if debug != best {
			fmt.Println(debug, best, joltage)
			panic("eff")
		}
	}

	memo[hash] = best
	return best
}

func Run() {
	defer AH.TrackTime(time.Now(), "Day 10")
	is, _ := AH.ReadStrFile("../inputs/day10.txt")
	p1, p2 := 0, 0
	for i, s := range is {
		j1, j2, moves := parseInput2(s)
		bm := binMap(moves)
		p1 += part1(j1, bm)
		// fmt.Println(j2, moves)
		memo = make(map[Hash]int)
		v := wtf2(j2, bm)
		p2 += v
		fmt.Println(i, p2, v)
	}
	AH.PrintSoln(10, p1, p2)

	return
}
