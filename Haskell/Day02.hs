
import Data.List.Split (splitOn)
import AdventHelper ( allEqual, printSoln, digitCount )

parseInput :: String -> [(Int, Int)]
parseInput s = map (\x -> (read (head x) :: Int, read (last x) :: Int)) t
    where t = map (splitOn "-") (splitOn "," s)

intBlock :: Int -> Int -> [Int]
intBlock 0 _ = []
intBlock n s = (n `mod` (10 ^ s)) : intBlock (n `div` (10 ^ s)) s
    
repDigit :: Int -> Int -> Bool
repDigit m n
    | digitCount m `mod` n == 0 = allEqual $ intBlock m s
    | otherwise = False
    where s = digitCount m `div` n

repDigit2 :: Int -> Bool
repDigit2 n = any (repDigit n) t
    where t = [ x | x <- [2..digitCount n], digitCount n `mod` x == 0]

goodInRange :: (Int, Int) -> [Int]
goodInRange (s,e) = filter (`repDigit` 2) [s..e]

goodInRange2 :: (Int, Int) -> [Int]
goodInRange2 (s,e) = filter repDigit2 [s..e]

main :: IO ()
main = do
    f <- readFile "../inputs/day02.txt"
    let ms = parseInput $ head $ lines f

    let p1 = sum $ map (sum . goodInRange) ms
    let p2 = sum $ map (sum . goodInRange2) ms
    printSoln 2 p1 p2
