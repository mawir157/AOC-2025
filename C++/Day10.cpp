#ifdef DAY10

#include "AH.h"

namespace Day10
{

    typedef std::vector<int64_t> Move;

    struct PreMove {
        Move delta;
        int64_t parity;
    };

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
    std::map<Hash, std::set<int64_t>> g_parity_memo;

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

    struct ScanResults
    {
        int64_t hi;
        int64_t lo;
        bool even;
    };

    ScanResults scan(const Joltage& j)
    {
        ScanResults s{INT64_MIN,INT64_MAX,true};

        for(auto x : j)
        {
            s.hi = std::max(s.hi,x);
            s.lo = std::min(s.lo,x);
            s.even &= !(x&1);
        }

        return s;
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

    void applyMove(Joltage & js, const Move & move, const bool part2)
    {
        for (size_t i = 0; i < js.size(); i++) {
            js[i] -= move[i];
            if (!part2) {
                js[i] %= 2;
            }
        }
        return;
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
        return std::popcount((uint64_t)n);
    }

    std::vector<PreMove> binMap(const std::vector<Move> & moves)
    {
        std::vector<PreMove> bm(1 << moves.size());
        for (int64_t bin = 0; bin < (1 << moves.size()); bin++) {
            auto bbin = bin;
            Move move(moves[0].size());
            for (size_t idx = 0; bbin > 0; idx++) {
                if (bbin & 1) {
                    auto mv = moves[idx];
                    applyMove(move, mv, true);
                }
                bbin >>= 1;
            }
            for (size_t idx = 0; idx < move.size(); idx++) {
                move[idx] *= -1;
            }
            PreMove pm = {move, countBits(bin)};
            bm[bin] = std::move(pm);
        }

        return bm;
    }

    std::set<int64_t>
    parity(const Joltage & js, const std::vector<PreMove> & bm, const bool part2)
    {
        auto hash = hashJoltage(js);
        if (g_parity_memo.count(hash) > 0) {
            return g_parity_memo[hash];
        }
        
        std::set<int64_t> valid_sequences;
        for (int64_t bin = 0; bin < (int64_t)bm.size(); bin++) {
            Joltage jtg = js;
            applyMove(jtg, bm[bin].delta, part2);
            auto [hi, lo, _] = scan(jtg);
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

        g_parity_memo[hash] = valid_sequences;
        return valid_sequences;
    }

    int64_t part1(const Joltage & js, const std::vector<PreMove> & bm)
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

    int64_t part2(const Joltage & joltage, const std::vector<PreMove> & bm, const bool first)
    {
        auto hash = hashJoltage(joltage);
        if (g_memo.count(hash) > 0) {
            return g_memo[hash];
        }

        auto [hi, lo, even] = scan(joltage);

    	if (hi == 0 && lo == 0) {
		    return 0;
	    }

        auto even_init = first && even;
        auto legal_moves = parity(joltage, bm, true);
        if (even_init) {
            for (int64_t bin = 0; bin < (int64_t)bm.size(); bin++) {
                legal_moves.insert(bin);
            }
        }

        auto best = INF;
        if (legal_moves.size() == 0) {
            g_memo[hash] = best;
            return best;
        }

        for (auto im : legal_moves) {
            Joltage jtg = joltage;
            applyMove(jtg, bm.at(im).delta, true);
            auto count = bm.at(im).parity;
            if (even_init) {
                applyMove(jtg, bm.at(im).delta, true);
                count += bm.at(im).parity;
            }
            // at this point next Joltage is all even
            auto [hhi, llo, _] = scan(jtg);

            if (llo < 0) {
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
            g_parity_memo.clear();
            p2 += part2(j2, bm, true);
        }

		AH::PrintSoln(10, p1, p2);

		return 0;
	}

}

#endif
