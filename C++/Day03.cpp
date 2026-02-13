#ifdef DAY03

#include "AH.h"

namespace Day03
{

	int64_t largestJoltage(std::string s, int64_t l) 
	{
		auto to_drop = s.size() - l;
		std::vector<int64_t> is;

		for (auto r : s) {
			std::string rs{r};
			int64_t i = AH::stoi64(rs);
			is.push_back(i);
		}

		int64_t jolt = 0;

		for (int i = 0; i + to_drop < is.size(); i++) {
			int64_t max_val = is[i];
			int64_t max_idx = 0;
			for (size_t j = 0; j < to_drop + 1; j++) {
				if (is[i+j] > max_val) {
					max_val = is[i+j];
					max_idx = j;
				}
			}

			jolt = 10*jolt + max_val;
			i += max_idx;
			to_drop -= max_idx;
		}

		return jolt;
	}
	
	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		int64_t part1 = 0;
		int64_t part2 = 0;
		for (auto s : ss) {
			part1 += largestJoltage(s, 2);
			part2 += largestJoltage(s, 12);
		}

		AH::PrintSoln(3, part1, part2);

		return 0;
	}

}

#endif
