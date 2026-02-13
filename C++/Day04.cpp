#ifdef DAY04

#include "AH.h"

namespace Day04
{

	typedef std::pair<int, int> Pos;
	
	std::map<Pos, int> parseInput(std::vector<std::string> ss)
	{
		std::map<Pos, int> m;
		
		for (size_t r = 0; r < ss.size(); r++) {
			auto s = ss[r];
			for (size_t c = 0; c < s.size(); c++) {
				auto ch = s.at(c);
				if (ch == '@') {
					Pos p{r,c};
					m[p] = 0;
				}
			}
		}

		return m;
	}

	std::vector<Pos> nbrs(Pos p)
	{
		return {
			Pos{p.first - 1, p.second - 1},
			Pos{p.first - 1, p.second},
			Pos{p.first - 1, p.second + 1},
			Pos{p.first, p.second - 1},
			Pos{p.first, p.second + 1},
			Pos{p.first + 1, p.second - 1},
			Pos{p.first + 1, p.second},
			Pos{p.first + 1, p.second + 1}
		};
	}

	void countNbrs(std::map<Pos, int> & m)
	{
		for (auto [p,c] : m) {
			auto ns = nbrs(p);
			m[p] = 0;
			for (auto n : ns) {
				m[p] += (m.count(n)) ? 1 : 0;
			}
		}
		
		return;
	}

	int clearWarehouse(std::map<Pos, int> & m)
	{
		int cleared = 0;
		std::vector<Pos> xx;

		while (true) {
			for (auto [p,n] : m) {
				if (n < 4) {
					xx.push_back(p);
				}
			}

			if (xx.size() == 0) {
				break;
			}

			cleared += xx.size();

			for (auto p : xx) {
				m.erase(p);
			}
			xx.clear();
			xx.reserve(m.size());
			countNbrs(m);
		}

		return cleared;
	}

	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto wh = parseInput(ss);
		countNbrs(wh);
		int64_t part1 = 0;
		for (auto [_,n]: wh) {
			part1 += (n < 4) ? 1 : 0;
		}
		int64_t part2 = clearWarehouse(wh);

		AH::PrintSoln(4, part1, part2);

		return 0;
	}

}

#endif
