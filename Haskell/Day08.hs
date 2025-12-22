import Data.List.Split (splitOn)
import AdventHelper (printSoln, thd)
import Data.List (sortBy, nub, sort, maximumBy, minimumBy)
import Data.Function (on)
import qualified Data.Set as Set

type Point = (Int, Int, Int)

parseInput :: String -> Point
parseInput s = (read x :: Int, read y :: Int, read z :: Int)
    where [x,y,z] = splitOn "," s

distSq :: Point -> Point -> Int
distSq (x,y,z) (s,r,t) = (x-s)*(x-s) + (y-r)*(y-r) + (z-t)*(z-t)

compPts :: Point -> Point -> Bool
compPts (x,y,z) (s,r,t)
    | x /= s = x < s
    | y /= r = y < r
    | z /= t = z < t
    | otherwise = False

pairDistances :: [Point] -> [((Point, Point), Int)]
pairDistances ps =  [((p, q), distSq p q) | p <- ps, q <- ps, compPts p q] 

getCC :: [Point] -> [(Point, Point)] -> [Point]
getCC [] _ = []
getCC vs [] = vs
getCC vs es
    | null vs' = vs
    | otherwise = getCC (nub (vs ++ vs')) xes
    where es' = filter (\(p,q) -> (p `elem` vs) || (q `elem` vs)) es
          vs' = nub (map fst es' ++ map snd es')
          xes = filter (`Set.notMember` Set.fromList es') es

allCCs :: [Point] -> [(Point, Point)] -> [[Point]]
allCCs [] _ = []
allCCs v es = getCC [head v] es : allCCs v' es
    where vcc = getCC [head v] es
          v' = filter (`Set.notMember` Set.fromList vcc) v

nearestNeighbour :: [Point] -> Point -> (Point, Point, Int)
nearestNeighbour ps p = minimumBy (compare `on` thd) ds
    where ds = [(p, q, distSq p q) | q <- ps, p /= q]

main :: IO ()
main = do
    f <- readFile "../inputs/day08.txt"
    let vs = map parseInput $ lines f
    let es' = map fst $ sortBy (compare `on` snd) $ pairDistances vs
    let p1 = product $ take 3 $ reverse $ sort $ map length $ allCCs vs $ take 1000 es'
    let ((x,_,_),(r,_,_),_) = maximumBy (compare `on` thd) $ map (nearestNeighbour vs) vs

    printSoln 8 p1 (x*r)
