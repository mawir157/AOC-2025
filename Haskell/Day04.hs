import qualified Data.Set as Set
import AdventHelper (printSoln)

parseInput :: (Int, Int) -> [String] -> Set.Set (Int, Int) -> Set.Set (Int, Int)
parseInput _ [] m = m
parseInput (r,c) ([]:ts) m = parseInput (r+1,0) ts m
parseInput (r,c) ((s:ss):ts) m
    | s == '@' = parseInput (r,c+1) (ss:ts) $ Set.insert (r,c) m
    | otherwise = parseInput (r,c+1) (ss:ts) m

nbrs :: (Int, Int) -> [(Int, Int)]
nbrs (r,c) = [(r-1,c-1),(r-1,c),(r-1,c+1),(r,c-1),(r,c+1),(r+1,c-1),(r+1,c),(r+1,c+1)]

countNbrs :: Set.Set (Int, Int) -> (Int, Int) -> Int
countNbrs s p = length $ filter (`Set.member` s) $ nbrs p

clearWarehouse :: Bool -> Set.Set (Int, Int) -> Set.Set (Int, Int)
clearWarehouse False s = s
clearWarehouse b s
    | null tk = clearWarehouse False s
    | otherwise = clearWarehouse True s'
    where tk = filter (\x -> countNbrs s x < 4) $ Set.toList s
          s' = foldr Set.delete s tk

main :: IO ()
main = do
    f <- readFile "../inputs/day04.txt"
    let ls = lines f
    let s = Set.empty
    let si = parseInput (0,0) ls s

    let p1 = length $ filter (< 4) $ map (countNbrs si) $ Set.toList si
    let p2 = length si - length (clearWarehouse True si)
    printSoln 4 p1 p2
