import Data.List.Split (splitOn)
import AdventHelper (printSoln)
import Data.List (sortBy)
import Data.Function (on)

parseInput :: String -> (Int, Int)
parseInput s = (read $ head t ::Int, read $ last t :: Int)
    where t= splitOn "-" s

merge :: [(Int, Int)] -> [(Int,Int)]
merge [x] = [x]
merge ((a,b):(y,z):xs)
    | b < y = (a,b) : merge ((y,z):xs)
    | otherwise = merge((a,max b z):xs)

inInterval :: [(Int, Int)] -> Int -> Bool
inInterval [] _ = False
inInterval ((x,y):xs) p = ((x <= p) && (p <= y)) || inInterval xs p

main :: IO ()
main = do
    f <- readFile "../inputs/day05.txt"
    let is = map parseInput $ takeWhile (not . null) $ lines f
    let is' = merge $ sortBy (compare `on` fst) is
    let ps = map (\x -> read x :: Int) $ drop 1 $ dropWhile (not. null ) $ lines f
    let p1 = length $ filter (inInterval is') ps
    let p2 = sum $ map (\(a,b) -> b - a + 1) is'

    printSoln 5 p1 p2
