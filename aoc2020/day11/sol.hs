import System.IO

main = do
    contents <- readFile "input_ex.txt"
    let seats = lines contents
    print(part1 seats)

loopSimulate :: [String] -> [String]
loopSimulate seats = 
    let (res, changed) = simulate [] seats
    in if changed
       then loopSimulate res
       else res

simulate :: String -> [String] -> ([String], Bool)
simulate _ [] = ([], False)
simulate prev seats =
    let (cur, next) = case (seats) of
            (c:[]) -> (c,"")
            (c:n:t) -> (c, n)
        (newCur, chng) = foldl (\(res, changed) posChar -> 
                             let (nChar, nChanged) = swap prev cur next posChar
                             in (res ++ [nChar], nChanged || changed)
                         ) ("", False) (enumerate cur)
        newPrev = cur
        (ending, cEnd) = simulate newPrev (tail seats)
    in ([newCur] ++ ending, chng || cEnd)

enumerate = zip [0..]

swap :: String -> String -> String -> (Int, Char) -> (Char, Bool)
swap prev cur next (pos, seat)
    | (seat == 'L') && ((countSquare pos prev cur next) == 0) = ('#', True)
    | (seat == '#') && ((countSquare pos prev cur next) >= 4) = ('L', True)
    | otherwise = (seat, False)


countSquare :: Int -> String -> String -> String -> Int
countSquare pos prev cur next =
    let lb = maximum [0, pos - 1]
        ub = minimum [(length cur-1), pos + 1]
        occupiedPrev = 
           if not (null prev)
           then foldl (\c s -> if prev!!s == '#' then c + 1 else c) 0 [lb..ub]
           else 0
        occupiedNext =
            if not (null next)
            then foldl (\c s -> if next!!s == '#' then c + 1 else c) 0 [lb..ub]
            else 0
        occupiedCur = foldl (\c s -> if (s /= pos) && (cur!!s) == '#' then c + 1 else c) 0 [lb..ub] 
    in occupiedPrev + occupiedCur + occupiedNext

countOccRow :: String -> Int
countOccRow = length . (filter (\c -> c == '#'))

countOccupied :: [String] -> Int
countOccupied seats = sum [ countOccRow row | row <- seats ]

part1 :: [String] -> Int
part1 seats = countOccupied (loopSimulate seats)
