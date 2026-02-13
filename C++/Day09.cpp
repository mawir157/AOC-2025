#ifdef DAY09

#include "AH.h"

namespace Day09
{

	typedef std::pair<int64_t, int64_t> Pos;
	struct Edge
	{
		Pos p, q;
		bool vert;
	};
	
	std::pair<std::vector<Pos>, std::vector<Edge>>
	parseInput(const std::vector<std::string> & ss)
	{
		std::vector<Pos> ps;

		for (auto s : ss) {
			auto pr = AH::Split(s, ',');
			ps.emplace_back(AH::stoi64(pr[0]), AH::stoi64(pr[1]));
		}

		std::vector<Edge> es;

		for (size_t i = 0; i < ps.size(); i++) {
			if (i == ps.size() - 1) {
				es.emplace_back(ps[i], ps[0], ps[i].first == ps[0].first);
			} else {
				es.emplace_back(ps[i], ps[i+1], ps[i].first == ps[i+1].first);
			}
		}

		return {ps, es};
	}

	bool goodRect(const Pos v1, const Pos v2, std::vector<Edge> es)
	{
		if ((v1.first == v2.first) || (v1.second == v2.second)) {
			return false;
		}

		bool good = true;

		for (auto e : es) {
			if (e.vert) {
				auto xLo = std::min(v1.first, v2.first);
				auto xHi = std::max(v1.first, v2.first);

				if ((xLo < e.p.first) && (e.p.first < xHi)) {
					for (int64_t y = e.p.second; y != e.q.second; y += AH::sgn(e.q.second - e.p.second)) {
						if ((y - v1.second) * (y - v2.second) < 0) {
							good = false;
							break;
						}
					}
				}
			} else {
				auto yLo = std::min(v1.second, v2.second);
				auto yHi = std::max(v1.second, v2.second);

				if ((yLo < e.p.second) && (e.p.second < yHi)) {
					for (int64_t x = e.p.first; x != e.q.first; x += AH::sgn(e.q.first - e.p.first)) {
						if ((x - v1.first) * (x - v2.first) < 0) {
							good = false;
							break;
						}
					}
				}			  
			}

			if (!good) {
				break;
			}
		}

		return good;
	}

	std::pair<int64_t, int64_t>
	maxRect(const std::vector<Pos> & ps, const std::vector<Edge> & es)
	{
		std::pair<int64_t, int64_t> areas{0,0};


		for (size_t iq = 0; iq < ps.size(); iq++) {
			auto q = ps[iq];
			for (size_t ip = iq + 1; ip < ps.size(); ip++) {
				auto p = ps[ip];
				auto area = (std::abs(q.first - p.first) + 1) * (std::abs(q.second - p.second) + 1);

				areas.first = std::max(area, areas.first);

				if ((area > areas.second) && goodRect(p, q, es)) {
					areas.second = area;
				}
			}
		}

		return areas;
	}
	
	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto [ps, es] = parseInput(ss);
		auto [part1, part2] = maxRect(ps, es);

		AH::PrintSoln(9, part1, part2);

		return 0;
	}

}

#endif
