package main

import (
	D01 "AoC2025/Day01"
	D02 "AoC2025/Day02"
	D03 "AoC2025/Day03"
	D04 "AoC2025/Day04"
	D05 "AoC2025/Day05"
	D06 "AoC2025/Day06"
	D07 "AoC2025/Day07"
	D08 "AoC2025/Day08"
	D09 "AoC2025/Day09"
	D10 "AoC2025/Day10"
	D11 "AoC2025/Day11"
	D12 "AoC2025/Day12"
	AH "AoC2025/adventhelper"
	"time"
)

func main() {
	defer AH.TrackTime(time.Now(), "All days")
	D01.Run()
	D02.Run()
	D03.Run()
	D04.Run()
	D05.Run()
	D06.Run()
	D07.Run()
	D08.Run()
	D09.Run()
	D10.Run()
	D11.Run()
	D12.Run()
}
