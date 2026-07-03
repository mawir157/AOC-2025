#ifdef DAY10

#include "AH.h"

namespace Day10
{

    typedef std::vector<int64_t> Move;
    typedef std::vector<int64_t> Joltage;

    struct Hash
    {
        int64_t q1, q2, q3;

        bool operator<(const Hash& other) const
        {
            return std::tie(q1, q2, q3) < std::tie(other.q1, other.q2, other.q3);
        }
    };

    int64_t INF = 1000000;
    
    std::map<Hash, int64_t> g_memo;

    int64_t moveSize(const Move & m)
    {
        int64_t sum = 0;
        for (auto v : m) {
            sum += v;
        }

        return sum;
    }

    Hash hashJoltage(const Joltage & js)
    {
        int64_t q1 = 0, q2 = 0, q3 = 0;
        for (size_t i = 0; i < js.size(); i++) {
            int64_t j = js[i];
            if (i < 4) {
                q1 *= 512;
                q1 += j;
            } else if (i < 8) {
                q2 *= 512;
                q2 += j;
            } else if (i < 12){
                q3 *= 512;
                q3 += j;
            }
        }

        return Hash{q1, q2, q3};
    }

    std::tuple<Joltage, Joltage, std::vector<Move>>
    parseInput(const std::string & s)
    {
        auto ss = AH::Split(s, ' ');
        size_t n = ss[0].size() - 2;

        auto t = ss[0].substr(1, ss[0].size() - 2);
        Joltage jolt1(t.size());
        for (size_t i = 0; i < t.size(); i++) {
            jolt1[i] = (t.at(i) == '#') ? 1 : 0;
        }

        std::vector<Move> moves;
        for (size_t i = 0; i < ss.size(); i++) {
            if ((i == 0) || (i == ss.size() - 1)) {
                continue;
            }
            auto sq = ss[i];
            Move move(n);
            t = sq.substr(1, s.size() - 1);
            auto ns = AH::Split(t, ',');
            for (auto n : ns) {
                move[AH::stoi64(n)] = 1;
            }
            moves.push_back(move);
        }

        Joltage jolt2;
        t = ss.back().substr(1, ss.back().size() - 1);
        auto ns = AH::Split(t, ',');
        for (auto n : ns) {
            jolt2.push_back(AH::stoi64(n));
        }

        return {jolt1, jolt2, moves};
    }

    Joltage applyMove(const Joltage & js, const Move & move, const bool part2)
    {
        auto js_copy = js;
        for (size_t i = 0; i < js.size(); i++) {
            if (part2) {
                js_copy[i] = js[i] - move[i];
            } else {
                js_copy[i] = (js[i] + move[i]) % 2;
            }
        }

        return js_copy;
    }

    bool allEven(const Joltage & js) {
        for (auto j : js) {
            if (j % 2 == 1) {
                return false;
            }
        }

        return true;
    }

    int64_t countBits(int64_t n)
    {
        int64_t bit_count = 0;
        while (n > 0) {
            if (n & 1) {
                bit_count++;
            }
            n /= 2;
        }

        return bit_count;
    }

    std::map<int64_t, Move> binMap(const std::vector<Move> & moves)
    {
        std::map<int64_t, Move> bm;
        for (int64_t bin = 0; bin < (1 << moves.size()); bin++) {
            auto bbin = bin;
            Move move(moves[0].size());
            for (size_t idx = 0; bbin > 0; idx++) {
                if (bbin & 1) {
                    auto mv = moves[idx];
                    move = applyMove(move, mv, true);
                }
                bbin /= 2;
            }
            for (size_t idx = 0; idx < move.size(); idx++) {
                move[idx] *= -1;
            }
            bm[bin] = move;
        }
        std::vector<int> to_kill;
        for (size_t bin1 = 0; bin1 < bm.size(); bin1++) {
            auto mv1 = bm[bin1];
            for (size_t bin2 = 0; bin2 < bm.size(); bin2++) {
                if (bin1 >= bin2) {
                    continue;
                }
                auto mv2 = bm[bin2];
                bool bad = true;
                for (size_t idx = 0; idx < mv2.size(); idx++) {
                    if (mv1[idx] != mv2[idx]) {
                        bad = false;
                        break;
                    }
                }
                if (bad) {
                    if (countBits(bin1) < countBits(bin2)) {
                        to_kill.push_back(bin2);
                    } else if (countBits(bin2) < countBits(bin1)) {
                        to_kill.push_back(bin1);
                    }                  
                }
            }
        }
        for (auto k : to_kill) {
            bm.erase(k);
        }

        return bm;
    }

    std::set<int64_t>
    parity(const Joltage & js, const std::map<int64_t, Move> bm, const bool part2)
    {
        std::set<int64_t> valid_sequences;
        for (auto [bin, mv] : bm) {
            auto jtg = applyMove(js, mv, part2);
            auto hi = *std::max_element(jtg.begin(), jtg.end());
            auto lo = *std::min_element(jtg.begin(), jtg.end());
            if (part2) {
                if ((allEven(jtg)) && (hi >= 0) && (lo >= 0)) {
                    valid_sequences.insert(bin);
                }
            } else {
                if ((hi == 0) && (lo == 0)) {
                    valid_sequences.insert(bin);
                }
            }
        }

        return valid_sequences;
    }

    int64_t part1(const Joltage & js, const std::map<int64_t, Move> bm)
    {
        int64_t val = INF;
        
        auto vs = parity(js, bm, false);
        for (auto v : vs) {
            auto vv = countBits(v);
            if (val > vv) {
                val = vv;
            }
        }

        return val;
    }

    int64_t part2(const Joltage & joltage, const std::map<int64_t, Move> bm, const bool first)
    {
        auto hash = hashJoltage(joltage);
        if (g_memo.count(hash) > 0) {
            return g_memo[hash];
        }

        auto hi = *std::max_element(joltage.begin(), joltage.end());
        auto lo = *std::min_element(joltage.begin(), joltage.end());

    	if (hi == 0 && lo == 0) {
		    return 0;
	    }

        auto even_init = first && allEven(joltage);
        auto legal_moves = parity(joltage, bm, true);
        if (even_init) {
            for (auto [k, _] : bm) {
                legal_moves.insert(k);
            }
        }

        auto best = INF;
        if (legal_moves.size() == 0) {
            g_memo[hash] = best;
            return best;
        }

        for (auto im : legal_moves) {
            Joltage jtg(joltage.begin(), joltage.end());
            jtg = applyMove(jtg, bm.at(im), true);
            auto count = countBits(im);
            if (even_init) {
                jtg = applyMove(jtg, bm.at(im), true);
                count += countBits(im);
            }
            // at this point next Joltage is all even
            auto hi = *std::max_element(jtg.begin(), jtg.end());
            auto lo = *std::min_element(jtg.begin(), jtg.end());

            if (lo < 0) {
                continue;
            }

            int64_t multiplier = 1;
            while (allEven(jtg) && hi > 0) {
                for (auto & j : jtg) {
                    j /= 2;
                }
                multiplier *= 2;
                break;
            }
            auto t = count + multiplier*part2(jtg, bm, false);
            if (best > t) {
                best = t;
            }
        }

        g_memo[hash] = best;
        return best;
    }

	int Run(const std::string& filename)
	{
		const auto ss = AH::ReadTextFile(filename);
        int64_t p1 = 0, p2 = 0;
        for (auto s : ss) {
            auto [j1, j2, moves] = parseInput(s);
            auto bm = binMap(moves);
            p1 += part1(j1, bm);
            g_memo.clear();
            p2 += part2(j2, bm, true);
        }

		AH::PrintSoln(10, p1, p2);

		return 0;
	}

}

#endif
