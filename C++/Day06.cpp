#ifdef DAY06

#include "AH.h"

namespace Day06
{

	std::pair<AH::Grid<uint64_t>, std::string> parseInput(const std::vector<std::string> ss)
	{
		AH::Grid<uint64_t> grid;
		std::string ops = "";

		for (uint64_t i = 0; i < ss.size(); i++) {
			auto s = ss[i];
			auto ps = AH::Fields(s);
			if (i != ss.size() - 1) {
				std::vector<uint64_t> row;
				for (auto p : ps) {
					auto n = AH::stoui64(p);
					row.emplace_back(n);
				}
				grid.push_back(row);
			} else {
				for (auto r : ps) {
					ops += r;
				}
			}
		}

		return {grid, ops};
	}

	uint64_t doMaths(const AH::Grid<uint64_t> & g, const std::string ops)
	{
		uint64_t total = 0;
		size_t rows = g.size();

		for (size_t col = 0; col < ops.size(); col++)
		{
			auto op = ops.at(col);
			uint64_t colt = (op == '*') ? 1 : 0;
			
			for (size_t row = 0; row < rows; row++) {
				if (op == '+') {
					colt += g[row][col];
				} else {
					colt *= g[row][col];
				}
			}
			total += colt;
		}

		return total;
	}

	uint64_t doCephalopodMaths(const std::vector<std::string> & ss)
	{
		size_t rows = ss.size();
		size_t cols = ss[0].size();

		auto op = ' ';
		auto new_sum = true;
		uint64_t big_sum = 0;
		uint64_t col_val = 0;

		for (size_t col = 0; col < cols; col++) {
			if (new_sum) {
				op = ss[rows-1].at(col);
				col_val = (op == '*') ? 1 : 0;
				new_sum = false;
			}
			uint64_t number = 0;
			for (size_t row = 0; row < rows-1; row++) {
				uint64_t digit = 0;
				auto ch = ss[row].at(col);
				if (ch != ' ') {
					number *= 10;
					digit = uint64_t(ch - '0');
				} 
				number += digit;
			}

			if (number == 0) {
				big_sum += col_val;
				new_sum = true;
			} else {
				if (op == '+') {
					col_val += number;
				} else {
					col_val *= number;
				}
			}

			if (col == cols-1) {
				big_sum += col_val;
			}
		}

		return big_sum;
	}

	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto [grid, ops] = parseInput(ss);
		auto part1 = doMaths(grid, ops);
		auto part2 = doCephalopodMaths(ss);

		AH::PrintSoln(6, part1, part2);

		return 0;
	}

}

#endif
