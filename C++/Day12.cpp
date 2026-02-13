#ifdef DAY12

#include "AH.h"

namespace Day12
{

	typedef std::array<std::array<bool, 3>, 3> Shape;
	struct Puzzle
	{
		uint64_t x, y;
		std::array<uint64_t, 6> counts;
	};
	
	std::pair<std::array<Shape, 6>, std::vector<Puzzle>>
	parseInput(const std::vector<std::string> & ss)
	{
		size_t shape_idx = 0;
		std::array<Shape, 6> shapes;
		for (shape_idx = 0; shape_idx < shapes.size(); shape_idx++) {
			for (size_t row_idx = 0; row_idx < 3; row_idx++) {
				auto s = ss[5*shape_idx + row_idx + 1];
				for (size_t col_idx = 0; col_idx < s.size(); col_idx++) {
					shapes[shape_idx][row_idx][col_idx] =  (s.at(col_idx) == '#');
				}
			}
		}

		std::vector<Puzzle> pzs;

		for (shape_idx = shape_idx * 5; shape_idx < ss.size(); shape_idx++) {
			Puzzle pz;
			auto ps = AH::SplitOnString(ss[shape_idx], ": ");
			auto dims = AH::Split(ps[0], 'x');
			pz.x = AH::stoui64(dims[0]);
			pz.y = AH::stoui64(dims[1]);

			auto cts = AH::Split(ps[1], ' ');
			for (size_t ni = 0; ni < cts.size(); ni++) {
				pz.counts[ni] = AH::stoui64(cts[ni]);
			}

			pzs.push_back(pz);
		}

		return {shapes, pzs};
	}

	bool solvePuzzle(const std::array<Shape, 6> & shapes, const Puzzle & p)
	{
		uint64_t bricks = 0;
		uint64_t tiles = 0;

		for (size_t i = 0; i < p.counts.size(); i++) {
			auto c = p.counts[i];
			for (size_t idx = 0; idx < 9; idx++) {
				bricks += (shapes[i][idx%3][idx/3]) ? c : 0;
			}
			tiles += c;
		}

		if (bricks > p.x * p.y) {
			return false;
		}

		if (tiles <= (p.x/3)*(p.y/3)) {
			return true;
		}

		return false; // never hit!
	}

	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto [pieces, puzzles] = parseInput(ss);
        uint64_t part1 = 0;
        std::string part2 = "You are important.";
        for (auto pz : puzzles) {
            part1 += (solvePuzzle(pieces, pz)) ? 1 : 0;
        }

		AH::PrintSoln(12, part1, part2);

		return 0;
	}

}

#endif
