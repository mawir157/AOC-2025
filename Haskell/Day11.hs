import Data.List.Split (splitOn)
import AdventHelper (printSoln)
import qualified Data.Map as Map
import Data.Maybe (fromJust)
import Data.List (foldl')

type Children = Map.Map String [String]
type Memo = Map.Map String Int

parseInput :: Children -> String -> Children
parseInput m s = Map.insert (head ps) (splitOn " " (last ps)) m
    where ps = splitOn ": " s

routes :: Children -> Memo ->String -> String -> (Memo, Int)
routes children memo t f
    | f == t = (memo, 1)
    | Just n <- Map.lookup f memo = (memo, n)
    | not $ Map.member f children = (Map.insert f 0 memo, 0)
    | otherwise = (Map.insert f total memo', total)
    where cs = fromJust $ Map.lookup f children
          (memo', total) = foldl' (\(acc_m, acc_n) c ->
              let (new_m, n) = routes children acc_m t c
              in (new_m, acc_n + n)) (memo, 0) cs

part2 :: Children -> String -> String -> Int
part2 m t f
    | mid1 /=0 = snd (routes m Map.empty "fft" f) * mid1 * snd (routes m Map.empty t "dac")
    | mid2 /=0 = snd (routes m Map.empty "dac" f) * mid2 * snd (routes m Map.empty t "fft")
    where mid1 = snd $ routes m Map.empty "dac" "fft"
          mid2 = snd $ routes m Map.empty"fft" "dac"          

main :: IO ()
main = do
    f <- readFile "../inputs/day11.txt"
    let ms = foldl parseInput Map.empty $ lines f
    -- let p1 = snd $ routes ms "out" "you"
    let memo = Map.empty :: Memo
    let p1 = snd $ routes ms Map.empty "out" "you"
    let p2 = part2 ms "out" "svr"

    printSoln 11 p1 p2
