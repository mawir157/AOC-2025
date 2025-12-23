import AdventHelper (printSoln)
import Data.List.Split (splitOn)

type Piece = [[Bool]]
type Puzzle = (Int, Int, [Int])

parsePieces :: [String] -> [Piece]
parsePieces [] = []
parsePieces (_:ss) = ps : parsePieces (drop 4 ss)
    where ps = map (map (== '#')) $ take 3 ss

parsePuzzles :: [String] -> [Puzzle]
parsePuzzles [] = []
parsePuzzles (s:ss) = (read $ head ds :: Int, read $ last ds :: Int, map read cs :: [Int]) : parsePuzzles ss
    where t = splitOn ": " s
          ds = splitOn "x" $ head t
          cs = splitOn " " $ last t

pieceSize :: Piece -> Int
pieceSize = foldr ((+) . length . filter id) 0

basicPuzzleChecks :: [Piece] -> Puzzle -> Maybe Bool
basicPuzzleChecks ps (x,y,cs)
    | (x*y) < tt = Just False
    | sum cs <= (x `div` 3) * (y `div` 3) = Just True
    | otherwise = Nothing
    where tt = sum $ zipWith (*) cs $ map pieceSize ps

main :: IO ()
main = do
    f <- readFile "../inputs/day12.txt"
    let ps = parsePieces $ take 30 $ lines f
    let pzs = parsePuzzles $ drop 30 $ lines f
    let p1 = length $ filter (== Just True) $ map (basicPuzzleChecks ps) pzs

    printSoln 12 p1 "Good bye. I love you."
