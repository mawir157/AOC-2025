import Data.Maybe (fromJust, isJust, isNothing)
import Data.List (elemIndex)
import AdventHelper (printSoln)
import qualified Data.Map as Map

parseInput :: Int -> String -> [Int]
parseInput _ [] = []
parseInput i (c:ss)
    | c == '^' = i : parseInput (i+1) ss
    | otherwise = parseInput (i+1) ss

doSplit :: (Int, Map.Map Int Int) -> [Int] -> (Int, Map.Map Int Int)
doSplit m [] = m
doSplit (a, m) (x:xs)
    | isNothing v = doSplit (a, m) xs
    | otherwise   = doSplit (a+ 1, m') xs
    where v = Map.lookup x m
          v' = fromJust v
          m' = Map.delete x $ Map.insertWith (+) (x+1) v' $ Map.insertWith (+) (x-1) v' m

main :: IO ()
main = do
    f <- readFile "../inputs/day07.txt"
    let ls = lines f
    let si = fromJust $ elemIndex 'S' $ head ls
    let os = map snd . filter ((==1) . fst) . zip (cycle [0,1]) $ drop 1 ls
    let ts = map (parseInput 0) os
    let m = Map.fromList [(si, 1)]

    let (p1, m') = foldl doSplit (0, m) ts 
    let p2 = sum $ map snd $ Map.toList m'

    printSoln 7 p1 p2
