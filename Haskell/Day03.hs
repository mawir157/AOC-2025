import AdventHelper (printSoln)
import Data.Char (digitToInt)
import Data.Maybe (fromJust)
import Data.List (elemIndex)

joltage :: Int -> Int -> [Int] -> Int
joltage a _ [] = a
joltage a 0 _ = a
joltage a l is = joltage (10*a + m) (l-1) (drop (idx' + 1) is)
    where m = maximum $ take (length is - l + 1) is
          idx' = fromJust $ elemIndex m is

main :: IO ()
main = do
    f <- readFile "../inputs/day03.txt"
    let ms = map (map digitToInt) $ lines f

    let p1 = sum $ map (joltage 0 2) ms
    let p2 = sum $ map (joltage 0 12) ms
    printSoln 3 p1 p2
