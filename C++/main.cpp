#define TIME

#include "days.h"
#include "AH.h"

int main()
{

	#ifdef DAY01
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day01::Run("../inputs/day01.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY02
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day02::Run("../inputs/day02.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY03
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day03::Run("../inputs/day03.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY04
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day04::Run("../inputs/day04.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY05
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day05::Run("../inputs/day05.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY06
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day06::Run("../inputs/day06.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY07
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day07::Run("../inputs/day07.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY08
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day08::Run("../inputs/day08.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY09
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day09::Run("../inputs/day09.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY10
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day10::Run("../inputs/day10.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY11
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day11::Run("../inputs/day11.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif

	#ifdef DAY12
		#ifdef TIME
		AH::start = std::chrono::steady_clock::now();
		#endif
	Day12::Run("../inputs/day12.txt");
		#ifdef TIME
		AH::end = std::chrono::steady_clock::now();
		AH::printTime();
		#endif
	#endif
	
	return 0;
}
