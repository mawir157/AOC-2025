#ifdef DAY05

#include "AH.h"

namespace Day05
{

	typedef std::pair<uint64_t, uint64_t> Interval;
	
	std::pair<std::vector<Interval>, std::vector<uint64_t>>
	parseInput(const std::vector<std::string> ss)
	{
		std::vector<Interval> rs;
		std::vector<uint64_t> ns;
		
		auto intervals = true;

		for (auto s : ss) {
			if (s.size() == 0) {
				intervals = false;
				continue;
			}
			if (intervals) {
				auto pr = AH::Split(s, '-');
				auto lhs = AH::stoui64(pr[0]);
				auto rhs = AH::stoui64(pr[1]);
				rs.emplace_back(lhs, rhs);
			} else {
				ns.push_back(AH::stoui64(s));
			}
		}

		return {rs, ns};
	}

	std::pair<uint64_t, uint64_t> checkIngredients(
		std::vector<Interval> & rs,
		std::vector<uint64_t> ns
	) 
	{
		std::pair<uint64_t, uint64_t> results;
		
		std::sort(rs.begin(), rs.end(), [](auto &left, auto &right) {
			return left.first < right.first;
		});

		std::vector<Interval> merged_rs;

		auto A = rs[0];
		for (size_t idx = 0; idx < rs.size(); idx++) {
			auto B = rs[idx];
			if (A.second < B.first) {
				merged_rs.push_back(A);
				A = B;
			} else {
				A.second = std::max(A.second, B.second);
			}
		}

		merged_rs.push_back(A);

		for (auto r : merged_rs) {
			results.second += r.second - r.first + 1;
		}

		for (auto n : ns) {
			for (auto i :  merged_rs) {
				if ((i.first <= n) && (n <= i.second)) {
					results.first++;
					break;
				}
			}
		}

		return results;
	}
	
	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto [rs, ns] = parseInput(ss);
		auto [part1, part2] = checkIngredients(rs, ns);

		AH::PrintSoln(5, part1, part2);

		return 0;
	}

}

#endif
