#ifdef DAY11

#include "AH.h"

namespace Day11
{

	typedef AH::Grid<bool> Adj;

	std::pair<AH::Grid<bool>, std::map<std::string, uint64_t>>
	buildGraph(const std::vector<std::string> & ss)
	{
		uint64_t label = 0;
		std::map<std::string, uint64_t> map_labels;

		for (auto s : ss) {
			auto ps = AH::SplitOnString(s, ": ");

			if (map_labels.count(ps[0]) == 0) {
				map_labels[ps[0]] = label;
				label++;
			}

			auto qs = AH::Split(ps[1], ' ');
			for (auto q : qs) {
				if (map_labels.count(q) == 0) {
					map_labels[q] = label;
					label++;
				}		
			}
		}
		AH::Grid<bool> adj(map_labels.size());
		for (auto & row : adj) {
			row = std::vector<bool>(map_labels.size()); 
		}

		for (auto s : ss) {
			auto ps = AH::SplitOnString(s, ": ");
			auto from = map_labels[ps[0]];

			auto qs = AH::Split(ps[1], ' ');
			for (auto q : qs) {
				auto to = map_labels[q];
				adj[from][to] = 1;
			}		
		}

		return {adj, map_labels};
	}

	uint64_t routes(
		const AH::Grid<bool> adj,
		const uint64_t from,
		const uint64_t to,
		std::map<std::pair<uint64_t, uint64_t>, uint64_t> & memo
	)
	{
		if (memo.count({from, to}) > 0) {
			return memo[{from, to}];
		} 

		if (from == to) {
			return 1;
		}

		uint64_t paths = 0;

		std::vector<uint64_t> children;
		for (size_t i = 0; i < adj.size(); i++) {
			if (adj[from][i]) {
				children.push_back(i);
			}
		}

		for (auto p : children) {
			paths += routes(adj, p, to, memo);
		}

		memo[{from, to}] = paths;
		return paths;
	}

	uint64_t routes2(
		const AH::Grid<bool> adj,
		const uint64_t from,
		const uint64_t to,
		const uint64_t fix1,
		const uint64_t fix2,
		std::map<std::pair<uint64_t, uint64_t>, uint64_t> & memo	
	)
	{
		uint64_t r = 1;

		auto mid1 = routes(adj, fix1, fix2, memo);
		auto mid2 = routes(adj, fix2, fix1, memo);

		if (mid1 != 0) {
			r = mid1;
			r *= routes(adj, from, fix1, memo);
			r *= routes(adj, fix2, to, memo);
		} else {
			r = mid2;
			r *= routes(adj, from, fix2, memo);
			r *= routes(adj, fix1, to, memo);
		}

		return r;
	}

	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto [adj, labels] = buildGraph(ss);
		std::map<std::pair<uint64_t, uint64_t>, uint64_t> cache;
		auto part1 = routes(adj, labels["you"], labels["out"], cache);
		auto part2 = routes2(adj, labels["svr"], labels["out"],
		                     labels["fft"], labels["dac"], cache);

		AH::PrintSoln(11, part1, part2);

		return 0;
	}

}

#endif
