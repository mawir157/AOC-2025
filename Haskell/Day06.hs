import AdventHelper (printSoln, if')

frontNumber :: String -> (Int, String)
frontNumber s = (n, ss)
    where n = read (takeWhile (/= ' ') s) :: Int
          ss = dropWhile (== ' ') $ dropWhile (/= ' ') s

calc :: Char -> [Int] -> Int
calc c = if' (c == '*') product sum

doNormalMaths :: (String, String, String, String, String) -> Int
doNormalMaths ([],_,_,_,_) = 0
doNormalMaths (s1,s2,s3,s4,c:sym) = calc c [n1,n2,n3,n4] + doNormalMaths (s1',s2',s3',s4',sym')
    where (n1,s1') = frontNumber s1
          (n2,s2') = frontNumber s2
          (n3,s3') = frontNumber s3
          (n4,s4') = frontNumber s4
          sym' = dropWhile (== ' ') sym

doCephalopodMaths :: ([Int], Char) -> (String, String, String, String, String) -> Int
doCephalopodMaths (ns,sym) ([],_,_,_,_) = calc sym ns
doCephalopodMaths (ns,sym) (x:xs,y:ys,z:zs,w:ws,c:cs)
    | null ns= doCephalopodMaths ([n], c) (xs,ys,zs,ws, cs)
    | null nn = calc sym ns + doCephalopodMaths ([], 'X') (xs,ys,zs,ws, cs)
    | otherwise = doCephalopodMaths (n:ns, sym) (xs,ys,zs,ws, cs)
    where nn = filter (/= ' ') [x,y,z,w]
          n = read nn :: Int

main :: IO ()
main = do
    f <- readFile "../inputs/day06.txt"
    let is = lines f
    let ss = (head is, is!!1, is!!2, is!!3, is!!4)

    let p1 = doNormalMaths ss
    let p2 = doCephalopodMaths ([], 'X') ss
    printSoln 6 p1 p2
