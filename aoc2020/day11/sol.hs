import System.IO

main = do
    contents <- readFile "input.txt"
    let seats = lines contents
    print(part1 seats)
    print(part2 seats)

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

-- Only goes right to go left use reverse row
searchRow :: String -> Int
searchRow [] = 0
searchRow (x:xs) = if x == '#' 
                   then 1 
                   else if x == 'L' then 0 else searchRow xs

-- Only goes down, to go up use reverse seats
searchCol :: Int -> [String] -> Int
searchCol _ [] = 0
searchCol pos (x:xs) = if insp == '#' 
                       then 1 
                       else if insp == 'L' then 0 else searchCol pos xs
                       where insp = x!!pos

-- Goes down right
-- - To go up right: use reverse seats
-- - To go down left: use (rowSize - 1 - pos) and (map reverse seats)
-- - To go up left: use reverse (map reverse seats) and rowSize - 1 - pos
searchDiag :: Int -> [String] -> Int
searchDiag _ [] = 0
searchDiag pos (x:xs)
    | (pos+1) >= (length x) = 0
    | otherwise =
        if insp == '#'
        then 1
        else if insp == 'L' then 0 else searchDiag (pos+1) xs
        where insp = x!!(pos+1)

countVisible :: (Int, Int) -> Int -> Int -> [String] -> [String] -> [String] -> [String] -> Int
countVisible (x,y) rowSize colSize seats urSeats dlSeats ulSeats =
    let curRow = seats !! x
        rCount = searchRow (drop (y+1) curRow)
        bottom = drop (x+1) seats
        drCount = searchDiag y bottom
        dCount = searchCol y bottom
        lBottom = drop (x+1) dlSeats
        dlCount = searchDiag (rowSize - 1 - y) lBottom
        lCurRow = dlSeats !! x
        lCount = searchRow (drop (rowSize - y) lCurRow)
        lTop = drop (colSize - x) ulSeats
        ulCount = searchDiag (rowSize - 1 - y) lTop
        rTop = drop (colSize - x) urSeats
        uCount = searchCol y rTop
        urCount = searchDiag y rTop
    in rCount + drCount + dCount + dlCount + lCount + ulCount + uCount + urCount

swapVisible :: ((Int, Int), Char) -> Int -> Int -> [String] -> [String] -> [String] -> [String] -> (Char, Bool)
swapVisible ((x,y), seat) rowSize colSize seats urSeats dlSeats ulSeats
    | (seat == 'L') && ((countVisible (x,y) rowSize colSize seats urSeats dlSeats ulSeats) == 0) = ('#', True)
    | (seat == '#') && ((countVisible (x,y) rowSize colSize seats urSeats dlSeats ulSeats) >= 5) = ('L', True)
    | otherwise = (seat, False)

simulateVisible :: Int -> Int -> [String] -> [String] -> [String] -> [String] -> ([String], Bool)
simulateVisible rowSize colSize seats urSeats dlSeats ulSeats =
    foldl (\(nSeats, sChng) (x, row) ->
        let (newRow, changed) = foldl (\(nRow, rChng) (y, seat) ->
               let (nSeat, chng) = swapVisible ((x,y), seat) rowSize colSize seats urSeats dlSeats ulSeats
               in (nRow ++ [nSeat], rChng || chng)
               ) ("", False) (enumerate row)
        in (nSeats ++ [newRow], sChng || changed)
    ) ([], False) (enumerate seats)

loopSimulateVisible :: [String] -> [String]
loopSimulateVisible seats = 
    let urSeats = reverse seats
        dlSeats = map reverse seats
        ulSeats = reverse dlSeats
        colSize = length(seats)
        rowSize = length(head seats)
        (res, changed) = simulateVisible rowSize colSize seats urSeats dlSeats ulSeats
    in if changed
       then loopSimulateVisible res
       else res

part1 :: [String] -> Int
part1 seats = countOccupied (loopSimulate seats)

part2 :: [String] -> Int
part2 seats = 
    countOccupied(loopSimulateVisible seats)
    --let urSeats = reverse seats
    --    dlSeats = map reverse seats
    --    ulSeats = reverse dlSeats
    --    colSize = length(seats)
    --    rowSize = length(head seats)
    --    (res, changed) = simulateVisible rowSize colSize seats urSeats dlSeats ulSeats
    --    urRes = reverse res
    --    dlRes = map reverse res
    --    ulRes = reverse dlRes
    --    colSize2 = length(res)
    --    rowSize2 = length(head res)
    --in simulateVisible rowSize2 colSize2 res urRes dlRes ulRes
    --print(countVisible (3,3) rowSize colSize seats urSeats dlSeats ulSeats)
