import Data.List.Split (splitOn)
import AdventHelper (printSoln)
import qualified Data.Map as Map
import Data.Maybe (fromJust, fromMaybe)
import Debug.Trace (traceShow)

type Children = Map.Map String (Int, [String])
type Memo = Map.Map (String, String) Int

parseInput :: Children -> String -> Children
parseInput m s = Map.insert (head ps) (0, splitOn " " (last ps)) m
    where ps = splitOn ": " s

routes :: Children -> String -> String -> Int
routes m t f
    | f == t = 1
    | not $ Map.member f m = 0
    | mm /= 0 = mm
    | otherwise = sum $ map (routes m t) cs
    where (mm, cs) = fromJust $ Map.lookup f m


-- adjust :: Ord k => (a -> a) -> k -> Map k a -> Map k a

-- routesMemo :: Children -> String -> String -> Children
-- routesMemo m t f
--     | f == t = m
--     | not $ Map.member f m = 0
--     | mm /= 0 = m
--     | otherwise = sum $ map (routes m t) cs
--     where (mm, cs) = fromJust $ Map.lookup f m

part2 :: Children -> String -> String -> Int
part2 m t f
    | mid1 /=0 = (routes m "fft" f) * mid1 * (routes m t "dac")
    | mid2 /=0 = (routes m "dac" f) * mid2 * (routes m t "fft")
    where mid1 = (traceShow "mid1") (routes m "dac" "fft")
          mid2 = (traceShow "mid2") (routes m "fft" "dac")

main :: IO ()
main = do
    f <- readFile "../inputs/day11.txt"
    let m = Map.empty
    let ms = foldl parseInput m $ lines f
    print ms
    let p1 = routes ms "out" "you"
    let p2 = 0--part2 ms "out" "svr"

    printSoln 11 p1 p2