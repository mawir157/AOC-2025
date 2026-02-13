#ifdef DAY01

#include "AH.h"

namespace Day01
{

	std::vector<int> parseInput(const std::vector<std::string> & ss)
	{
		std::vector<int> moves;
		for (auto s : ss) {
			auto c = s.at(0);
			auto n = s.substr(1);
			auto nn = std::stoi(n);

			moves.push_back((c == 'R') ? nn : -1*nn);
		}
		return moves;
	}

	std::pair<int, int> combination(int pos, int vals, std::vector<int> ms) {
		int score1 = 0;
		int score2 = 0;

		for (auto m : ms) {
			auto s = pos;
			auto e = pos + m;

			score1 += (e % vals == 0) ? 1 : 0;
			if (s < e) {
				e /= vals;
				s /= vals;
			} else {
				e = (e-1) / vals;
				s = (s-1) / vals;
			}
			score2 += std::abs(e - s);
			pos += m;
		}

		return {score1, score2};
	}

	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
		auto ms = parseInput(ss);
		auto [part1, part2] = combination(50, 100, ms);

		AH::PrintSoln(1, part1, part2);

		return 0;
	}

}

#endif
