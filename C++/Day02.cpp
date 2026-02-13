// #ifdef DAY02

#include "AH.h"

namespace Day02
{

	std::vector<std::pair<uint64_t, uint64_t>> parseInput(const std::string & s)
	{
		auto ss = AH::Split(s, ',');
		std::vector<std::pair<uint64_t, uint64_t>> ps;

		for (auto p : ss) {
			auto pr = AH::Split(p, '-');
			ps.emplace_back(AH::stoui64(pr[0]), AH::stoui64(pr[1]));
		}

		return ps;
	}

	int countDigits(uint64_t i) {
		int digits = 0;
		while (i > 0) {
			digits++;
			i /= 10;
		}

		return digits;
	}
	
	std::pair<uint64_t, uint64_t> countInvalid(std::pair<uint64_t, uint64_t> p) {
		std::pair<uint64_t, uint64_t> counts(0,0);

		for (uint64_t i = p.first; i <= p.second; i++) {
			auto d = countDigits(i);
			for (int reps = 2; reps <= d; reps++) {
				if (d%reps != 0) {
					continue;
				}
				uint64_t ii = i;
				uint64_t mask = AH::IntPow(10, d/reps);
				auto is_rep_unit = true;
				uint64_t rep_block = ii % mask;
				ii /= mask;
				while (ii > 0) {
					if (ii%mask != rep_block) {
						is_rep_unit = false;
						break;
					}
					ii /= mask;
				}

				if (is_rep_unit) {
					if (reps == 2) {
						counts.first += i;
					}
					counts.second += i;
					break;
				}
			}
		}

		return counts;
	}

	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto ms = parseInput(ss[0]);
		uint64_t part1 = 0;
		uint64_t part2 = 0;
		for (auto m : ms) {
			auto [p1, p2] = countInvalid(m);
			part1 += p1;
			part2 += p2;
		}

		AH::PrintSoln(2, part1, part2);

		return 0;
	}

}

// #endif
