#ifdef DAY07

#include "AH.h"

namespace Day07
{
	
	std::pair<uint64_t, AH::Grid<uint64_t>> parseInput(std::vector<std::string> ss)
	{
		AH::Grid<uint64_t> ts;
		size_t start_idx = 0;

		for (size_t r = 0; r < ss.size(); r++) {
			if (r%2 == 1) {
				continue;
			}
			std::vector<uint64_t> row;
			auto s = ss[r];
			for (size_t idx = 0; idx < s.size(); idx++) {
				auto c = s.at(idx);
				if (c == '^') {
					row.push_back(idx);
				}
				if (c == 'S') {
					start_idx = idx;
				}
			}
			if (r > 0) {
				ts.push_back(row);
			}
		}

		return {start_idx, ts};
	}

	std::pair<uint64_t, uint64_t> beamSplitter(uint64_t s, AH::Grid<uint64_t> g)
	{
		uint64_t splits = 0;
		std::map<uint64_t, uint64_t> beam;
		beam[s] = 1;

		for (auto r : g) {
			for (auto [b,_] : beam) {
				for (auto sp : r) {
					if (b == sp) {
						splits++;
						beam[b-1] += beam[b];
						beam[b+1] += beam[b];
						beam.erase(b);
						break;
					}
				}
			}
		}

		uint64_t timelines = 0;
		for (auto [_,v] :  beam) {
			timelines += v;
		}

		return {splits, timelines};
	}
	
	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto [s,g] = parseInput(ss);
		auto [part1, part2] = beamSplitter(s,g);

		AH::PrintSoln(7, part1, part2);

		return 0;
	}

}

#endif
