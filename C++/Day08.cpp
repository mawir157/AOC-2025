#ifdef DAY08

#include "AH.h"

namespace Day08
{

	struct Pos
	{
		uint64_t x,y,z;
	};
	
	std::vector<Pos> parseInput(const std::vector<std::string> & ss)
	{
		std::vector<Pos> ps;

		for (auto s : ss) {
			auto pr = AH::Split(s, ',');
			ps.push_back(Pos{
				AH::stoui64(pr[0]),
				AH::stoui64(pr[1]),
				AH::stoui64(pr[2])
			});
		}

		return ps;
	}

	AH::Grid<bool> buildGraph(const std::vector<Pos> & ps)
	{
		AH::Grid<bool> adj(ps.size());
		for (size_t i = 0; i < ps.size(); i++) {
			adj[i] = std::vector<bool>(ps.size());
		}

		size_t pIdx = 0;
		size_t qIdx = 0;

		for (size_t i = 0; i < ps.size(); i++) {
			uint64_t dist = 1000000000000;
			for (size_t iq = 0; iq < ps.size(); iq++) {
				auto q = ps[iq];
				for (size_t ip = iq + 1; ip < ps.size(); ip++) {
					auto p = ps[ip];
					if (adj[iq][ip]) {
						continue;
					}

					uint64_t dist_sqrd =
						(p.x-q.x)*(p.x-q.x) +
						(p.y-q.y)*(p.y-q.y) +
						(p.z-q.z)*(p.z-q.z);
					
					if (dist_sqrd <  dist) {
						pIdx = ip;
						qIdx = iq;
						dist = dist_sqrd;
					}
				}
			}
			adj[qIdx][pIdx] = true;
			adj[pIdx][qIdx] = true;
		}

		return adj;
	}

	uint64_t findConComps(const AH::Grid<bool> & adj)
	{
		AH::Grid<uint64_t> comps;
		std::vector<bool> flagged(adj.size());
		uint64_t count = adj.size();
		uint64_t open = adj.size() - 1;

		while (count > 0) {
			std::vector<uint64_t> comp;
			count = 0;
			
			std::set<int> queue;
			queue.insert(open);
			flagged[open] = true;
			comp.push_back(open);

			while (queue.size() > 0) {
				uint64_t p = 0;
				for (auto k : queue) {
					p = k;
					break;
				}
				queue.erase(p);
				for (size_t q = 0; q < adj[p].size(); q++) {
					auto b = adj[p][q];
					if (b && !flagged[q]) {
						comp.push_back(q);
						flagged[q] = true;
						queue.insert(q);
					}
				}
			}

			for (size_t i = 0; i < flagged.size(); i++) {
				if (!flagged[i]) {
					count++;
					open = i;
				}
			}
			comps.push_back(comp);
		}

		std::vector<uint64_t> cs;
		for (auto c : comps) {
			cs.push_back(c.size());
		}

		std::sort(cs.begin(), cs.end(), std::greater<>());
		return cs[0] * cs[1] * cs[2];
	}

	uint64_t mostIsolatedVector(std::vector<Pos> ps)
	{
		std::vector<uint64_t> nns(ps.size());
		std::vector<size_t> nps(ps.size());

		for (size_t ip = 0; ip < ps.size(); ip++) {
			auto p = ps[ip];
			uint64_t nn = 1000000000000;
			for (size_t iq = 0; iq < ps.size(); iq++) {
				if (iq == ip) {
					continue;
				}
				auto q = ps[iq];
				auto distSqrd =
					(p.x-q.x)*(p.x-q.x) +
					(p.y-q.y)*(p.y-q.y) +
					(p.z-q.z)*(p.z-q.z);
				if (distSqrd < nn) {
					nn = distSqrd;
					nps[ip] = iq;
				}				
			}
			nns[ip] = nn;
		}

		uint64_t max = nns[0];
		size_t max_idx = 0;

		for (size_t i = 0; i < nns.size(); i++) {
			if (max < nns[i]) {
				max = nns[i];
				max_idx = i;
			}
		}

		return ps[max_idx].x * ps[nps[max_idx]].x;
	}

	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto ps = parseInput(ss);
		auto adj = buildGraph(ps);
		auto part1 = findConComps(adj);
		auto part2 = mostIsolatedVector(ps);

		AH::PrintSoln(8, part1, part2);

		return 0;
	}

}

#endif
