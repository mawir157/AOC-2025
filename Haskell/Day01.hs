import AdventHelper

parseInput :: String -> Int
parseInput s = if' (head s == 'R') 1 (-1) * (read (tail s) :: Int)

f1 :: Int -> (Int, Int, Int) -> Int -> (Int, Int, Int)
f1 v (ptr, p1, p2) m = (ptr', p1 + inc1, p2 + inc2)
    where ptr' = ptr + m 
          inc1 = if' (ptr' `mod` v == 0) 1 0
          inc2 = length $ filter (\x -> x `mod` v == 0) $ init [ptr,ptr+signum (ptr' - ptr)..ptr']

main :: IO ()
main = do
    f <- readFile "../inputs/day01.txt"
    let ms = map parseInput $ lines f
    let (_, p1, p2) = foldl (f1 100) (50, 0, 0) ms 

    printSoln 1 p1 p2
